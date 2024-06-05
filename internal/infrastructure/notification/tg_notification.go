package notification

import (
	"awesomeProject/internal/infrastructure/database/postgres/database"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/crypto/bcrypt"
)

// TelegramBot представляет собой структуру для работы с ботом Telegram.
type TelegramBot struct {
	bot         *tgbotapi.BotAPI
	chatID      int64
	db          *database.PGDatabase
	stopChannel chan struct{}
}

// NewTelegramBot создает новый экземпляр TelegramBot.
func NewTelegramBot(token string, chatID string, db *database.PGDatabase) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("error creating Telegram bot: %v", err)
	}

	chatIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		log.Fatalf("error parsing chat ID: %v", err)
	}

	return &TelegramBot{
		bot:         bot,
		chatID:      chatIDInt,
		db:          db,
		stopChannel: make(chan struct{}),
	}
}

// StartListening начинает прослушивание сообщений бота.
func (bot *TelegramBot) StartListening() {
	log.Println("Telegram bot is now listening for messages")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.bot.GetUpdatesChan(u)
	if err != nil {
		log.Printf("error getting updates channel: %v", err)
		return
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		message := update.Message
		// Handle incoming messages here
		// For example, to handle authorization via email:
		if strings.HasPrefix(message.Text, "/authorize") {
			args := strings.TrimPrefix(message.Text, "/authorize ")
			parts := strings.Split(args, " ")
			if len(parts) != 2 {
				bot.sendMessage(message.Chat.ID, "Usage: /authorize <email> <password>")
				continue
			}

			email := parts[0]
			password := parts[1]

			err := bot.authorizeUserByEmail(email, password, message.Chat.ID)
			if err != nil {
				log.Printf("error authorizing user: %v", err)
				bot.sendMessage(message.Chat.ID, "Error authorizing user. Please try again later.")
				continue
			}
			bot.sendMessage(message.Chat.ID, "User authorized successfully!")
		}
	}
}

func (bot *TelegramBot) authorizeUserByEmail(email, password string, chatID int64) error {
	ctx := context.Background()
	var userID int64
	var storedPassword string

	err := bot.db.QueryRow(ctx, "SELECT id, password FROM users WHERE email = $1", email).Scan(&userID, &storedPassword)
	if err != nil {
		return fmt.Errorf("error getting user by email: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}

	// Логика авторизации, например, сохранение авторизованного пользователя в базе данных или предоставление доступа к определенным функциям.
	bot.NotifyBirthdaySubscribers(userID)

	return nil
}

func (bot *TelegramBot) NotifyBirthdaySubscribers(userID int64) {
	ctx := context.Background()
	now := time.Now()

	rows, err := bot.db.Query(ctx, "SELECT users.id, users.name, users.birthday FROM users JOIN subscriptions ON users.id = subscriptions.user_id WHERE subscriptions.subscribed_to_id = $1", userID)
	if err != nil {
		log.Println("error getting subscribers: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var subscriberID int64
		var subscriberName string
		var subscriberBirthday time.Time
		err := rows.Scan(&subscriberID, &subscriberName, &subscriberBirthday)
		if err != nil {
			log.Println("error scanning subscriber row: %v", err)
			continue
		}

		if isBirthday(now, subscriberBirthday) {
			log.Println(now, " ===============    ", subscriberBirthday)
			message := fmt.Sprintf("Сегодня день рождения у %s! 🎉", subscriberName)
			bot.sendMessage(bot.chatID, message)
		}
	}
}

// sendMessage отправляет сообщение через бота Telegram.
func (bot *TelegramBot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.bot.Send(msg)
	if err != nil {
		log.Println("error sending message: %v", err)
	}
}

// isBirthday проверяет, является ли указанная дата днем рождения.
func isBirthday(today, birthday time.Time) bool {
	log.Println(today.Day(), "=====", birthday.Day(), "=====", today.Month(), "=====", birthday.Month())
	return today.Day() == birthday.Day() && today.Month() == birthday.Month()
}

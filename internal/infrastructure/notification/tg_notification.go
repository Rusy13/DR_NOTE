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
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"golang.org/x/crypto/bcrypt"
)

// TelegramBot представляет собой структуру для работы с ботом Telegram.
type TelegramBot struct {
	bot         *tgbotapi.BotAPI
	chatID      int64
	db          *database.PGDatabase
	stopChannel chan struct{}
	apiID       int
	apiHash     string
	phone       string
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

	err := bot.db.QueryRow(ctx, "SELECT id, password, api_id, api_hash, phone FROM users WHERE email = $1", email).Scan(&userID, &storedPassword, &bot.apiID, &bot.apiHash, &bot.phone)
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

	var subscribers []int64
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
			subscribers = append(subscribers, subscriberID)
		}
	}

	if len(subscribers) > 0 {
		err = bot.createTelegramChat(ctx, bot.apiID, bot.apiHash, bot.phone, subscribers)
		if err != nil {
			log.Println("error creating Telegram chat: %v", err)
		}
	}
}

func (bot *TelegramBot) createTelegramChat(ctx context.Context, apiID int, apiHash, phone string, userIDs []int64) error {
	client := telegram.NewClient(apiID, apiHash, telegram.Options{})

	return client.Run(ctx, func(ctx context.Context) error {
		sendCodeOptions := auth.SendCodeOptions{
			AllowFlashCall: false,
			CurrentNumber:  false,
			AllowAppHash:   false,
		}
		sentCode, err := client.Auth().SendCode(ctx, phone, sendCodeOptions)
		if err != nil {
			return fmt.Errorf("ошибка отправки кода: %w", err)
		}

		sentCodeConcrete, ok := sentCode.(*tg.AuthSentCode)
		if !ok {
			return fmt.Errorf("неожиданный тип ответа от SendCode")
		}

		// Ask the user to input the code they received
		fmt.Print("Введите код, отправленный на ваш телефон: ")
		var code string
		if _, err := fmt.Scan(&code); err != nil {
			return err
		}

		// Sign in using the code
		auth, err := client.Auth().SignIn(ctx, phone, code, sentCodeConcrete.PhoneCodeHash)
		if err != nil {
			return fmt.Errorf("неправильный код")
		}
		fmt.Printf("Авторизация прошла успешно: %+v\n", auth.User)

		var inputUsers []tg.InputUserClass
		for _, userID := range userIDs {
			inputUser := &tg.InputUser{
				UserID:     int64(userID),
				AccessHash: 0,
			}
			inputUsers = append(inputUsers, inputUser)
		}

		cchat, err := tg.NewClient(client).MessagesCreateChat(ctx, &tg.MessagesCreateChatRequest{
			Users: inputUsers,
			Title: "My very normal title",
		})
		if err != nil {
			return fmt.Errorf("ошибка при создании чата: %w", err)
		}

		for _, userID := range inputUsers {
			soob, err := tg.NewClient(client).MessagesAddChatUser(ctx, &tg.MessagesAddChatUserRequest{ChatID: int64(cchat.TypeID()), UserID: userID})
			log.Println("soob.GetMissingInvitees():", soob.GetMissingInvitees())
			log.Println("err:", err)
		}

		return nil
	})
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

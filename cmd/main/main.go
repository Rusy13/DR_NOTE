package main

import (
	"awesomeProject/internal/infrastructure/notification"
	"context"
	"fmt"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"log"
	"net/http"

	"awesomeProject/internal/infrastructure/database/postgres/database"
	"awesomeProject/internal/infrastructure/database/redis"
	"awesomeProject/internal/middleware"
	"awesomeProject/internal/routes"
	"awesomeProject/internal/user/delivery"
	"awesomeProject/internal/user/service"
	"awesomeProject/internal/user/storage/database"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	apiID   = 22686505                           // Replace with your API ID
	apiHash = "8db223c8889c9e9cef33ee1746b039ad" // Replace with your API hash
	phone   = "+79818557454"                     // Replace with your phone number
)

func main() {
	client := telegram.NewClient(apiID, apiHash, telegram.Options{})

	if err := client.Run(context.Background(), func(ctx context.Context) error {
		// Send the code to the phone number
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

		// Создаем новый чат с указанными пользователями.
		userIDs := []int64{564241094} // Идентификаторы пользователей
		var inputUsers []tg.InputUserClass

		for _, userID := range userIDs {

			inputUser := &tg.InputUser{
				UserID:     int64(userID),
				AccessHash: 0, // Оставьте как 0 для обычных пользователей
			}

			inputUsers = append(inputUsers, inputUser)
			log.Println("Добавлен пользователь:", userID)
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
			log.Println("errrrrrrrrrrrrr:", err)
			log.Println("Добавлен пользователь:", userID)
		}

		return nil
	}); err != nil {
		log.Fatalf("ошибка при запуске клиента Telegram: %v", err)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		log.Println("Error is-----------------------", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error in logger initialization: %v", err)
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in logger sync: %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbPool, err := database.New(ctx)
	if err != nil {
		logger.Fatalf("error in database init: %v", err)
	}
	defer func() {
		err = dbPool.Close()
		if err != nil {
			logger.Errorf("error in closing db")
		}
	}()

	redisConn, err := redis.Init()
	if err != nil {
		logger.Fatalf("error on connection to redis: %v", err)
	}
	defer func() {
		err = redisConn.Close()
		if err != nil {
			logger.Infof("error on redis close: %s", err.Error())
		}
	}()

	stOrder := storage.New(dbPool, redisConn, logger)
	svOrder := service.New(stOrder, zapLogger)
	d := delivery.New(svOrder, logger)

	mw := middleware.New(logger)
	router := routes.GetRouter(d, mw)

	port := "8000"
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	go func() {
		logger.Fatal(http.ListenAndServe(addr, router))
	}()

	bot := notification.NewTelegramBot("6598556806:AAGLxlf-WDYRC0ZjaIqDVEAaQ8zS-PFT_hs", "716615282", dbPool)
	//TOKEN CHAT_ID
	go bot.StartListening()

	select {}

}

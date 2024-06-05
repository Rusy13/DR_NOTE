package main

import (
	"awesomeProject/internal/infrastructure/notification"
	"context"
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

func main() {
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

	bot := notification.NewTelegramBot("TOKEN", "CHAT_ID", dbPool)
	//TOKEN CHAT_ID
	go bot.StartListening()

	select {}

}

package service

import (
	"awesomeProject/internal/user/model"
	"awesomeProject/internal/user/storage"
	"context"
	"go.uber.org/zap"
)

type UserService interface {
	AddUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	GetSubscribers(ctx context.Context, subscribedToID int) ([]int, error)
	Subscribe(ctx context.Context, userID, subscribedToID int) error
	Unsubscribe(ctx context.Context, userID, subscribedToID int) error
}

type UserServiceApp struct {
	storage storage.Storage
	logger  *zap.Logger
}

func New(storage storage.Storage, logger *zap.Logger) *UserServiceApp {
	return &UserServiceApp{
		storage: storage,
		logger:  logger,
	}
}

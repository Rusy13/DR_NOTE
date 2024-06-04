package storage

import (
	"context"

	"awesomeProject/internal/user/model"
)

type Storage interface {
	AddUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	GetSubscribers(ctx context.Context, subscribedToID int) ([]int, error)
	Subscribe(ctx context.Context, userID, subscribedToID int) error
	Unsubscribe(ctx context.Context, userID, subscribedToID int) error
}

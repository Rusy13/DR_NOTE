package service

import (
	"awesomeProject/internal/user/model"
	"context"
	"log"
)

func (s *UserServiceApp) AddUser(ctx context.Context, user *model.User) (*model.User, error) {
	log.Println("Adding user", user)
	addedUser, err := s.storage.AddUser(ctx, user)
	if err != nil {
		log.Println("Failed to add user", err)
		return nil, err
	}
	return addedUser, nil
}

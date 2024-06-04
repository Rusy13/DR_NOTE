package service

import (
	"awesomeProject/internal/user/model"
	"context"
	"log"
)

func (s *UserServiceApp) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	log.Println("Getting user by ID", id)
	user, err := s.storage.GetUserByID(ctx, id)
	if err != nil {
		log.Println("Failed to get user by ID", err)
		return nil, err
	}
	return user, nil
}

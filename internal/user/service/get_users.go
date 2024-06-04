package service

import (
	"awesomeProject/internal/user/model"
	"context"
	"log"
)

func (s *UserServiceApp) GetUsers(ctx context.Context) ([]model.User, error) {
	s.logger.Info("Getting all users")
	users, err := s.storage.GetUsers(ctx)
	if err != nil {
		log.Println("Failed to get users", err)
		return nil, err
	}
	return users, nil
}

package service

import (
	"awesomeProject/internal/user/model"
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (s *UserServiceApp) AddUser(ctx context.Context, user *model.User) (*model.User, error) {
	log.Println("Adding user", user)

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password", err)
		return nil, err
	}
	user.Password = string(hashedPassword)

	addedUser, err := s.storage.AddUser(ctx, user)
	if err != nil {
		log.Println("Failed to add user", err)
		return nil, err
	}
	return addedUser, nil
}

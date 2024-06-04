package service

import (
	"context"
	"log"
)

func (s *UserServiceApp) Subscribe(ctx context.Context, userID, subscribedToID int) error {
	log.Println("Subscribing user ", "userID ", userID, "subscribedToID", subscribedToID)
	err := s.storage.Subscribe(ctx, userID, subscribedToID)
	if err != nil {
		log.Println("Failed to subscribe user", err)
		return err
	}
	return nil
}

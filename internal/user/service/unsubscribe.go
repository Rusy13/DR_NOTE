package service

import (
	"context"
	"log"
)

func (s *UserServiceApp) Unsubscribe(ctx context.Context, userID, subscribedToID int) error {
	log.Println("Unsubscribing user", "userID ", userID, "subscribedToID", subscribedToID)
	err := s.storage.Unsubscribe(ctx, userID, subscribedToID)
	if err != nil {
		log.Println("Failed to unsubscribe user", err)
		return err
	}
	return nil
}

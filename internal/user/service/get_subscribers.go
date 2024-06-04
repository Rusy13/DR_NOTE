package service

import (
	"context"
	"log"
)

func (s *UserServiceApp) GetSubscribers(ctx context.Context, subscribedToID int) ([]int, error) {
	log.Println("Getting subscribers", subscribedToID)
	subscribers, err := s.storage.GetSubscribers(ctx, subscribedToID)
	if err != nil {
		log.Println("Failed to get subscribers", err)
		return nil, err
	}
	return subscribers, nil
}

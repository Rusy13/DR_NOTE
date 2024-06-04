package storage

import (
	"context"
	"fmt"
)

func (s *UserStorageDB) Subscribe(ctx context.Context, userID, subscribedToID int) error {
	query := "INSERT INTO subscriptions (user_id, subscribed_to_id) VALUES ($1, $2)"
	_, err := s.db.Exec(ctx, query, userID, subscribedToID)
	if err != nil {
		s.logger.Errorw("Failed to subscribe user", "error", err)
		return err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("%s%d", subsCachePrefix, subscribedToID)
	s.redisConn.Do("DEL", cacheKey)

	return nil
}

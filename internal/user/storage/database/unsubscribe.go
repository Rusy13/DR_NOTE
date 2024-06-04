package storage

import (
	"context"
	"fmt"
)

func (s *UserStorageDB) Unsubscribe(ctx context.Context, userID, subscribedToID int) error {
	query := "DELETE FROM subscriptions WHERE user_id = $1 AND subscribed_to_id = $2"
	_, err := s.db.Exec(ctx, query, userID, subscribedToID)
	if err != nil {
		s.logger.Errorw("Failed to unsubscribe user", "error", err)
		return err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("%s%d", subsCachePrefix, subscribedToID)
	s.redisConn.Do("DEL", cacheKey)

	return nil
}

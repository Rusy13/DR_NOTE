package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func (s *UserStorageDB) GetSubscribers(ctx context.Context, subscribedToID int) ([]int, error) {
	cacheKey := fmt.Sprintf("%s%d", subsCachePrefix, subscribedToID)
	var subscribers []int

	// Try to get subscribers from cache
	cachedSubs, err := redis.Bytes(s.redisConn.Do("GET", cacheKey))
	if err == nil {
		if err := json.Unmarshal(cachedSubs, &subscribers); err == nil {
			return subscribers, nil
		}
	}

	// If not in cache, get subscribers from DB
	query := "SELECT user_id FROM subscriptions WHERE subscribed_to_id = $1"
	err = s.db.Select(ctx, &subscribers, query, subscribedToID)
	if err != nil {
		s.logger.Errorw("Failed to get subscribers from database", "error", err)
		return nil, err
	}

	// Cache the result
	subsBytes, _ := json.Marshal(subscribers)
	s.redisConn.Do("SETEX", cacheKey, s.cacheExpireTime, subsBytes)

	return subscribers, nil
}

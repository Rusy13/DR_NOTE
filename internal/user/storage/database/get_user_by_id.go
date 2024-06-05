package storage

import (
	models "awesomeProject/internal/user/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func (s *UserStorageDB) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	cacheKey := fmt.Sprintf("%s%d", userCachePrefix, id)
	user := &models.User{}

	// Try to get user from cache
	cachedUser, err := redis.Bytes(s.redisConn.Do("GET", cacheKey))
	if err == nil {
		if err := json.Unmarshal(cachedUser, user); err == nil {
			return user, nil
		}
	}

	// If not in cache, get user from DB
	query := "SELECT id, name, email, birthday, api_id, api_hash, phone FROM users WHERE id = $1"
	err = s.db.Get(ctx, user, query, id)
	if err != nil {
		s.logger.Errorw("Failed to get user by ID from database", "error", err)
		return nil, err
	}

	// Cache the result
	userBytes, _ := json.Marshal(user)
	s.redisConn.Do("SETEX", cacheKey, s.cacheExpireTime, userBytes)

	return user, nil
}

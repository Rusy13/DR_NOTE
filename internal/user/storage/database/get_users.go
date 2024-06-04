package storage

import (
	models "awesomeProject/internal/user/model"
	"context"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

func (s *UserStorageDB) GetUsers(ctx context.Context) ([]models.User, error) {
	cacheKey := usersCachePrefix
	var users []models.User

	// Try to get users from cache
	cachedUsers, err := redis.Bytes(s.redisConn.Do("GET", cacheKey))
	if err == nil {
		if err := json.Unmarshal(cachedUsers, &users); err == nil {
			return users, nil
		}
	}

	// If not in cache, get users from DB
	query := "SELECT id, name, email, birthday FROM users"
	err = s.db.Select(ctx, &users, query)
	if err != nil {
		s.logger.Errorw("Failed to get users from database", "error", err)
		return nil, err
	}

	// Cache the result
	usersBytes, _ := json.Marshal(users)
	s.redisConn.Do("SETEX", cacheKey, s.cacheExpireTime, usersBytes)

	return users, nil
}

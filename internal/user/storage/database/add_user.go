package storage

import (
	models "awesomeProject/internal/user/model"
	"context"
)

func (s *UserStorageDB) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := "INSERT INTO users (name, email, birthday, password, api_id, api_hash, phone) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := s.db.QueryRow(ctx, query, user.Name, user.Email, user.Birthday, user.Password, user.ApiID, user.ApiHash, user.Phone).Scan(&user.ID)
	if err != nil {
		s.logger.Errorw("Failed to add user to database", "error", err)
		return nil, err
	}

	// Invalidate cache
	s.redisConn.Do("DEL", usersCachePrefix)

	return user, nil
}

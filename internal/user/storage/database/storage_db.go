package storage

import (
	"awesomeProject/internal/infrastructure/database/postgres/database"
	"go.uber.org/zap"

	"github.com/gomodule/redigo/redis"
)

const (
	expireTime       = 15 * 60 // 15 минут в секундах
	userCachePrefix  = "user:"
	usersCachePrefix = "users:"
	subsCachePrefix  = "subs:"
)

type UserStorageDB struct {
	db              database.Database
	redisConn       redis.Conn
	cacheExpireTime int
	logger          *zap.SugaredLogger
}

func New(db database.Database, redisConn redis.Conn, logger *zap.SugaredLogger) *UserStorageDB {
	return &UserStorageDB{
		db:              db,
		logger:          logger,
		redisConn:       redisConn,
		cacheExpireTime: expireTime,
	}
}

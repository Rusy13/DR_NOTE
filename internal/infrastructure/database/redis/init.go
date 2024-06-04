package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Init() (redis.Conn, error) {
	//host := os.Getenv("REDIS_HOST")
	//port := os.Getenv("REDIS_PORT")
	host := "localhost"
	port := "6379"
	c, err := redis.DialURL(fmt.Sprintf("redis://user:@%s:%s/0", host, port))
	if err != nil {
		return nil, err
	}
	return c, nil

}

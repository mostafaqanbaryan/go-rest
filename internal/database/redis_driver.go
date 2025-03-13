package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisDriver() *redis.Client {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if rdb == nil {
		panic("NewRedisDriver failed")
	}
	if err := rdb.Ping(ctx); err != nil {
		panic(fmt.Errorf("NewRedisDriver %v", err))
	}
	return rdb
}

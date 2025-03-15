package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewRedisDriver() *redis.Client {
	ctx := context.Background()
	cacheHost := os.Getenv("REDIS_HOST")
	cachePassword := os.Getenv("REDIS_PASSWORD")

	cacheDatabase, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		cacheDatabase = 0
	}

	cachePort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		cacheDatabase = 6379
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cacheHost, cachePort),
		Password: cachePassword, // no password set
		DB:       cacheDatabase, // use default DB
	})

	if rdb == nil {
		panic("NewRedisDriver failed")
	}

	if err := rdb.Ping(ctx); err != nil {
		panic(fmt.Errorf("NewRedisDriver %v", err))
	}

	return rdb
}

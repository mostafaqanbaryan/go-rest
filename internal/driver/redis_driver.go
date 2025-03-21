package driver

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
)

type RedisDriver struct {
	redis *redis.Client
}

func NewRedisDriver() RedisDriver {
	ctx := context.Background()
	cacheHost := os.Getenv("REDIS_HOST")
	cachePassword := os.Getenv("REDIS_PASSWORD")

	cacheDatabase, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		cacheDatabase = 0
	}

	cachePort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		cachePort = 6379
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cacheHost, cachePort),
		Password: cachePassword, // no password set
		DB:       cacheDatabase, // use default DB
	})

	if rdb == nil {
		panic("NewRedisDriver failed")
	}

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Errorf("NewRedisDriver %v", err))
	}

	return RedisDriver{
		redis: rdb,
	}
}

func (d RedisDriver) Get(ctx context.Context, key string) (string, error) {
	get := d.redis.Get(ctx, key)
	if get == nil {
		return "", driverErrors.ErrGetCommand
	}

	result, err := get.Result()
	if err == redis.Nil {
		return "", driverErrors.ErrRecordNotFound
	} else if err != nil {
		return "", err
	}

	return result, nil
}

func (d RedisDriver) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	set := d.redis.Set(ctx, key, value, ttl)
	if set == nil {
		return driverErrors.ErrSetCommand
	}

	if err := set.Err(); err != nil {
		return err
	}
	return nil
}

func (d RedisDriver) Close() error {
	return d.redis.Close()
}

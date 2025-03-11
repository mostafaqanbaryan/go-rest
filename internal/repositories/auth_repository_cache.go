package repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type AuthRepositoryCache struct {
	db *redis.Client
}

func NewAuthRepositoryCache(rdb *redis.Client) AuthRepositoryCache {
	return AuthRepositoryCache{
		db: rdb,
	}
}

func (r AuthRepositoryCache) NewUserSession(user entities.User) (string, error) {
	sessionID := "1232"
	ctx := context.Background()
	_, err := r.db.Get(ctx, sessionID).Result()
	for {
		if err == redis.Nil {
			err := r.db.Set(ctx, "key", "value", 0).Err()
			if err != nil {
				return "", err
			}
			return sessionID, nil
		}
	}
}

package repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type AuthRepositoryCache struct {
	db *redis.Client
}

func NewAuthRepositoryCache(db *redis.Client) AuthRepositoryCache {
	return AuthRepositoryCache{
		db: db,
	}
}

func (r AuthRepositoryCache) NewUserSession(user entities.User) (string, error) {
	ctx := context.Background()
	for {
		sessionID := "1232"
		_, err := r.db.Get(ctx, sessionID).Result()
		if err == redis.Nil {
			err := r.db.Set(ctx, "key", "value", 0).Err()
			if err != nil {
				return "", err
			}
			return sessionID, nil
		}
	}
}

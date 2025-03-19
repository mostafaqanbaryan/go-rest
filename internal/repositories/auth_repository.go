package repositories

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"mostafaqanbaryan.com/go-rest/internal/database"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type CacheDriver interface {
	Set(context.Context, string, any, time.Duration) error
	Get(context.Context, string) (string, error)
}

type AuthRepositoryCache struct {
	db CacheDriver
}

func NewAuthRepository(db CacheDriver) AuthRepositoryCache {
	return AuthRepositoryCache{
		db: db,
	}
}

func (r AuthRepositoryCache) NewUserSession(user entities.User) (string, error) {
	ctx := context.Background()
	for {
		sessionID := generateSessionID()
		_, err := r.db.Get(ctx, sessionID)
		if err == database.ErrRecordNotFound {
			err := r.db.Set(ctx, sessionID, user.ID, time.Hour*10)
			if err != nil {
				return "", err
			}
			return sessionID, nil
		}
	}
}

func (r AuthRepositoryCache) GetUserIDBySessionID(sessionID string) (string, error) {
	ctx := context.Background()
	return r.db.Get(ctx, sessionID)
}

func generateSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

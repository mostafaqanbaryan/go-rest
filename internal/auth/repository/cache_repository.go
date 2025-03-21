package repository

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"

	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type cacheDriver interface {
	Set(context.Context, string, any, time.Duration) error
	Get(context.Context, string) (string, error)
}

type authRepository struct {
	db cacheDriver
}

func NewAuthRepository(db cacheDriver) authRepository {
	return authRepository{
		db: db,
	}
}

func (r authRepository) NewUserSession(user entities.User) (string, error) {
	ctx := context.Background()
	for {
		sessionID := generateSessionID()
		_, err := r.db.Get(ctx, sessionID)
		if err == driverErrors.ErrRecordNotFound {
			err := r.db.Set(ctx, sessionID, user.ID, time.Hour*10)
			if err != nil {
				return "", err
			}
			return sessionID, nil
		}
	}
}

func (r authRepository) GetUserIDBySessionID(sessionID string) (int64, error) {
	ctx := context.Background()
	res, err := r.db.Get(ctx, sessionID)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(res, 10, 0)
}

func generateSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

package authrepository

import (
	"context"
	"strconv"
	"time"

	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/pkg/strings"
)

type cacheDriver interface {
	Set(context.Context, string, any, time.Duration) error
	Get(context.Context, string) (string, error)
	Has(context.Context, string) (bool, error)
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
		sessionID := strings.GenerateRandom(32)
		exists, err := r.db.Has(ctx, sessionID)
		if err != nil {
			return "", err
		}

		if !exists {
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

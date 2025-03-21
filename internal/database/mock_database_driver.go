package database

import (
	"context"
	"database/sql"
	"math/rand"

	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type MockDatabaseDriver struct {
	list map[int64]any
}

func NewMockDatabaseDriver() MockDatabaseDriver {
	return MockDatabaseDriver{
		list: make(map[int64]any, 0),
	}
}

func (d MockDatabaseDriver) FindAllUsers(ctx context.Context) ([]entities.User, error) {
	list := make([]entities.User, len(d.list))
	for _, u := range d.list {
		list = append(list, u.(entities.User))
	}
	return list, nil
}

func (d MockDatabaseDriver) FindUserByUsername(ctx context.Context, username string) (entities.User, error) {
	for _, row := range d.list {
		user := row.(entities.User)
		if username == user.Username {
			return user, nil
		}
	}

	return entities.User{}, ErrRecordNotFound
}

func (d MockDatabaseDriver) CreateUser(ctx context.Context, params entities.CreateUserParams) (sql.Result, error) {
	userID := rand.Int63()
	d.list[userID] = entities.User{
		ID:       userID,
		Username: params.Username,
		Password: params.Password,
	}
	return nil, nil
}

func (d MockDatabaseDriver) UpdateUser(ctx context.Context, params entities.UpdateUserParams) error {
	d.list[params.ID] = params
	return nil
}

func (d MockDatabaseDriver) DeleteUser(ctx context.Context, userID int64) error {
	return nil
}

package repositories

import (
	"context"
	"database/sql"

	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/errors"
)

type DB interface {
	FindAllUsers(context.Context) ([]entities.User, error)
	FindUserByUsername(context.Context, string) (entities.User, error)
	CreateUser(context.Context, entities.CreateUserParams) (sql.Result, error)
	UpdateUser(context.Context, entities.UpdateUserParams) error
	DeleteUser(context.Context, int64) error
}

type UserRepositoryDB struct {
	db DB
}

func NewUserRepositoryDB(db DB) UserRepositoryDB {
	return UserRepositoryDB{
		db,
	}
}

func (r UserRepositoryDB) FindByUsername(username string) (entities.User, error) {
	ctx := context.Background()
	user, err := r.db.FindUserByUsername(ctx, username)
	switch err {
	case errors.ErrNotFound:
		return entities.User{}, errors.ErrUserNotFound
	}
	return user, nil
}

package repository

import (
	"context"
	"database/sql"

	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type DB interface {
	FindAllUsers(context.Context) ([]entities.User, error)
	FindUserByUsername(context.Context, string) (entities.User, error)
	FindUser(context.Context, int64) (entities.User, error)
	CreateUser(context.Context, entities.CreateUserParams) (sql.Result, error)
	UpdateUser(context.Context, entities.UpdateUserParams) error
	DeleteUser(context.Context, int64) error
}

type UserRepository struct {
	db DB
}

func NewUserRepository(db DB) UserRepository {
	return UserRepository{
		db,
	}
}

func (r UserRepository) FindByUsername(username string) (entities.User, error) {
	ctx := context.Background()
	user, err := r.db.FindUserByUsername(ctx, username)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
func (r UserRepository) FindUser(userID int64) (entities.User, error) {
	ctx := context.Background()
	user, err := r.db.FindUser(ctx, userID)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r UserRepository) Create(username, password string) error {
	ctx := context.Background()
	params := entities.CreateUserParams{
		Username: username,
		Password: password,
	}
	_, err := r.db.CreateUser(ctx, params)
	return err
}

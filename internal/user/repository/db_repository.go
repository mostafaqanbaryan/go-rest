package repository

import (
	"context"
	"database/sql"

	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type DB interface {
	FindAllUsers(context.Context) ([]entities.User, error)
	FindUserByEmail(context.Context, string) (entities.User, error)
	FindUser(context.Context, int64) (entities.User, error)
	CreateUser(context.Context, entities.CreateUserParams) (sql.Result, error)
	UpdatePassword(context.Context, entities.UpdatePasswordParams) error
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

func (r UserRepository) FindByEmail(email string) (entities.User, error) {
	ctx := context.Background()
	user, err := r.db.FindUserByEmail(ctx, email)
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

func (r UserRepository) Create(hashID, email, password string) error {
	ctx := context.Background()
	params := entities.CreateUserParams{
		HashID:   hashID,
		Email:    email,
		Password: password,
	}
	_, err := r.db.CreateUser(ctx, params)
	return err
}

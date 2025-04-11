package userrepository

import (
	"context"
	"database/sql"

	"mostafaqanbaryan.com/go-rest/internal/entities"
)

type sqlcConnection interface {
	FindAllUsers(context.Context) ([]entities.User, error)
	FindUserByEmail(context.Context, string) (entities.User, error)
	FindUser(context.Context, int64) (entities.User, error)
	CreateUser(context.Context, entities.CreateUserParams) (sql.Result, error)
	UpdateUser(context.Context, entities.UpdateUserParams) error
	DeleteUser(context.Context, int64) error
}

type userRepository struct {
	conn sqlcConnection
	db   *sql.DB
}

func NewUserRepository(db *sql.DB) userRepository {
	conn := entities.New(db)
	return userRepository{
		conn: conn,
		db:   db,
	}
}

func (r userRepository) FindByEmail(email string) (entities.User, error) {
	ctx := context.Background()
	user, err := r.conn.FindUserByEmail(ctx, email)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r userRepository) IsDuplicateEmail(email string) (bool, error) {
	ctx := context.Background()
	user, err := r.conn.FindUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (r userRepository) FindUser(userID int64) (entities.User, error) {
	ctx := context.Background()
	user, err := r.conn.FindUser(ctx, userID)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r userRepository) Create(hashID, email, password string) error {
	ctx := context.Background()
	params := entities.CreateUserParams{
		HashID:   hashID,
		Email:    email,
		Password: password,
	}
	_, err := r.conn.CreateUser(ctx, params)
	return err
}

func (r userRepository) Update(userID int64, fullname string) error {
	ctx := context.Background()
	params := entities.UpdateUserParams{
		ID:       userID,
		Fullname: sql.NullString{String: fullname},
	}
	err := r.conn.UpdateUser(ctx, params)
	return err
}

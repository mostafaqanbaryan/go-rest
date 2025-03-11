package repositories

import (
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/errors"
)

type Db interface {
	SelectAll(string, ...any) ([]any, error)
	SelectOne(string, ...any) (any, error)
}

type UserRepositoryDB struct {
	db Db
}

func NewUserRepositoryDB(db Db) UserRepositoryDB {
	return UserRepositoryDB{
		db: db,
	}
}

func (r UserRepositoryDB) FindByUsername(username string) (*entities.User, error) {
	if r.db == nil {
		return nil, nil
	}

	user, err := r.db.SelectOne("SELECT * FROM users WHERE username = ? LIMIT 1", username)
	switch err {
	case errors.ErrNotFound:
		return nil, errors.ErrUserNotFound
	}
	return user.(*entities.User), nil
}

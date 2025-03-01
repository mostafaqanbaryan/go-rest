package repositories

import "mostafaqanbaryan.com/go-rest/internal/entities"

type UserRepositoryDB struct {
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{}
}

func (r UserRepositoryDB) FindByUsername(username string) (*entities.User, error) {
	return nil, nil
}

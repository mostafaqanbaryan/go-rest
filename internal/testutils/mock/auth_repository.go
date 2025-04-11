package mock

import (
	"errors"

	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/pkg/strings"
)

type MockAuthRepository struct {
	List map[string]int64
}

func (r MockAuthRepository) NewUserSession(user entities.User) (string, error) {
	for {
		sessionID := strings.GenerateRandom(32)
		_, ok := r.List[sessionID]
		if ok {
			continue
		}

		r.List[sessionID] = user.ID
		return sessionID, nil
	}
}

func (r MockAuthRepository) GetUserIDBySessionID(sessionID string) (int64, error) {
	res, ok := r.List[sessionID]
	if !ok {
		return 0, errors.New("record not found")
	}
	return res, nil
}

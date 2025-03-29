package authservice

import "mostafaqanbaryan.com/go-rest/internal/entities"

type authRepository interface {
	NewUserSession(entities.User) (string, error)
	GetUserIDBySessionID(string) (int64, error)
}

type authService struct {
	repo authRepository
}

func NewAuthService(repo authRepository) authService {
	return authService{
		repo: repo,
	}
}

func (s authService) CreateSession(user entities.User) (string, error) {
	sessionId, err := s.repo.NewUserSession(user)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func (s authService) GetSession(sessionID string) (int64, error) {
	userID, err := s.repo.GetUserIDBySessionID(sessionID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

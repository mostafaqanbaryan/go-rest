package services

import "mostafaqanbaryan.com/go-rest/internal/entities"

type AuthRepository interface {
	NewUserSession(*entities.User) (string, error)
}

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s AuthService) CreateSession(user *entities.User) (string, error) {
	sessionId, err := s.repo.NewUserSession(user)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

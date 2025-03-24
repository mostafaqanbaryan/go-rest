package service_test

import (
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/auth/service"
	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/pkg/strings"
)

type mockAuthRepository struct {
	list map[string]int64
}

func (r mockAuthRepository) NewUserSession(user entities.User) (string, error) {
	for {
		sessionID := strings.GenerateRandom(32)
		_, ok := r.list[sessionID]
		if ok {
			continue
		}

		r.list[sessionID] = user.ID
		return sessionID, nil
	}
}

func (r mockAuthRepository) GetUserIDBySessionID(sessionID string) (int64, error) {
	res, ok := r.list[sessionID]
	if !ok {
		return 0, driverErrors.ErrRecordNotFound
	}
	return res, nil
}
func TestAuthService(t *testing.T) {
	t.Parallel()

	user := entities.User{
		ID:       123,
		Email:    "test",
		Password: "test",
	}

	authRepository := &mockAuthRepository{
		list: map[string]int64{},
	}

	authService := service.NewAuthService(authRepository)

	t.Run("Create session", func(t *testing.T) {
		sessionId, err := authService.CreateSession(user)
		if err != nil {
			t.Fatalf("want no error, got: <%v>", err)
		}

		userID, err := authService.GetSession(sessionId)
		if err != nil {
			t.Fatalf("want no error, got: <%v>", err)
		}

		if userID != user.ID {
			t.Fatalf("want email %d, got: %v", user.ID, userID)
		}
	})
}

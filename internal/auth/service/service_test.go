package service_test

import (
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/auth/repository"
	"mostafaqanbaryan.com/go-rest/internal/auth/service"
	"mostafaqanbaryan.com/go-rest/internal/driver"
	"mostafaqanbaryan.com/go-rest/internal/entities"
)

func TestAuthService(t *testing.T) {
	t.Parallel()

	user := entities.User{
		ID:       123,
		Username: "test",
		Password: "test",
	}

	cache := driver.NewMockCacheDriver()
	authRepository := repository.NewAuthRepository(cache)
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
			t.Fatalf("want username %d, got: %v", user.ID, userID)
		}
	})
}

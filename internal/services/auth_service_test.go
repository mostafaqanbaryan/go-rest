package services

import (
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/database"
	"mostafaqanbaryan.com/go-rest/internal/entities"
	"mostafaqanbaryan.com/go-rest/internal/repositories"
)

func TestAuthService(t *testing.T) {
	user := entities.User{
		ID:       123,
		Username: "test",
		Password: "test",
	}

	cache := database.NewMockCacheDriver()
	authRepository := repositories.NewAuthRepository(cache)
	authService := NewAuthService(authRepository)

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

package repositories

import (
	"testing"

	"mostafaqanbaryan.com/go-rest/internal/database"
)

func TestUserRepositoryDB(t *testing.T) {
	username := "test"
	t.Run("UserRepository find by username", func(t *testing.T) {
		db := database.NewMockDriver()
		rep := NewUserRepositoryDB(db)
		_, err := rep.FindByUsername(username)
		if err != nil {
			t.Fatalf("want %s, got %v", username, err)
		}
	})
}

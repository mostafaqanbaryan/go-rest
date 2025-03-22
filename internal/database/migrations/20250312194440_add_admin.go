package migrations

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pressly/goose/v3"
	"mostafaqanbaryan.com/go-rest/internal/argon2"
)

func init() {
	goose.AddMigrationContext(Up20250312194440, Down20250312194440)
}

func Up20250312194440(ctx context.Context, tx *sql.Tx) error {
	encrypted, err := argon2.CreateHash("123")
	if err != nil {
		return err
	}

	hashID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO users (hash_id, fullname, email, password) VALUES (?, 'admin', 'admin@go-rest', ?);
	`, hashID, encrypted)
	if err != nil {
		return err
	}
	return nil
}

func Down20250312194440(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DELETE FROM users WHERE email = 'admin@go-rest';`)
	return err
}

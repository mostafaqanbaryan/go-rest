package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func MigrateUp(db *sql.DB) {
	instance, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(fmt.Errorf("WithInstance: %w", err))
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../internal/database/migrations", "mysql", instance)
	if err != nil {
		panic(fmt.Errorf("NewWithDatabaseInstance: %w", err))
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Errorf("UpMigrate: %w", err))
	}
}

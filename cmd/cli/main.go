package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	_ "mostafaqanbaryan.com/go-rest/internal/database/migrations"
	"mostafaqanbaryan.com/go-rest/internal/driver"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pressly/goose/v3"
)

func main() {
	db := driver.NewMySQLDriver("")
	defer db.Close()

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	ctx := context.Background()
	err := goose.RunContext(ctx, "up", db, "internal/database/migrations")
	if err != nil {
		panic(err)
	}
}

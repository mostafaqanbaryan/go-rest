package main

import (
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "mostafaqanbaryan.com/go-rest/internal/database/migrations"
	"mostafaqanbaryan.com/go-rest/internal/driver"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pressly/goose/v3"
)

func main() {
	var action string
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "migrate":
		if len(os.Args) <= 2 {
			panic("missing command")
		}

		command := os.Args[2]

		db := driver.NewMySQLDriver("")
		defer db.Close()

		if err := goose.SetDialect("mysql"); err != nil {
			panic(err)
		}

		arguments := []string{}
		if len(os.Args) > 3 {
			arguments = append(arguments, os.Args[3:]...)
		}

		ctx := context.Background()
		err := goose.RunContext(ctx, command, db, "internal/database/migrations", arguments...)
		if err != nil {
			panic(err)
		}
	default:
		panic("unknown action")
	}
}

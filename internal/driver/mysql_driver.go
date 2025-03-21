package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDriver(connection string) *sql.DB {
	if connection == "" {
		dbName := os.Getenv("MYSQL_DATABASE")
		dbUser := os.Getenv("MYSQL_USER")
		dbPass := os.Getenv("MYSQL_PASSWORD")
		dbHost := os.Getenv("MYSQL_HOST")
		connection = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)
	}

	conn, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	if err = conn.Ping(); err != nil {
		panic(err)
	}
	return conn
}

package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDriver(connection string) *sql.DB {
	if connection == "" {
		connection = "root:root@/gorest?parseTime=true"
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

package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDriver() *sql.DB {
	conn, err := sql.Open("mysql", "root:root@/database?parseTime=true")
	if err != nil {
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err)
	}
	return conn
}

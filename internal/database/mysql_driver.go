package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDriver struct {
	conn *sql.DB
}

func (d MySQLDriver) SelectOne(query string, args ...any) (any, error) {
	rows, err := d.SelectAll(query, args)
	return rows[0], err
}

func (d MySQLDriver) SelectAll(query string, args ...any) ([]any, error) {
	statement, err := d.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(args)
	if err != nil {
		return nil, err
	}

	users := make([]any, 1)
	for rows.Next() {
		var temp any
		if err := rows.Scan(&temp); err != nil {
			return nil, err
		}
		users = append(users, temp)
	}
	return users, nil
}

func NewMySQLDriver() (*MySQLDriver, error) {
	conn, err := sql.Open("mysql", "root:root@/database")
	if err != nil {
		return nil, err
	}
	return &MySQLDriver{conn}, nil
}

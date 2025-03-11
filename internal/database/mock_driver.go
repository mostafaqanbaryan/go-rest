package database

import "mostafaqanbaryan.com/go-rest/internal/errors"

type MockDriver struct{}

func (d MockDriver) SelectOne(query string, args ...any) (any, error) {
	rows, err := d.SelectAll(query, args)
	return rows[0], err
}

func (MockDriver) SelectAll(string, ...any) ([]any, error) {
	return nil, errors.ErrNotFound
}

func NewMockDriver() MockDriver {
	return MockDriver{}
}

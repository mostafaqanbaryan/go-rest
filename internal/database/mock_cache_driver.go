package database

import (
	"context"
	"time"
)

type MockCacheDriver struct {
	list map[string]any
}

func NewMockCacheDriver() MockCacheDriver {
	return MockCacheDriver{
		list: make(map[string]any, 0),
	}
}

func (d MockCacheDriver) Get(ctx context.Context, key string) (string, error) {
	res, ok := d.list[key]
	if !ok {
		return "", ErrRecordNotFound
	}

	return res.(string), nil
}

func (d MockCacheDriver) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	d.list[key] = value
	return nil
}

func (d MockCacheDriver) Close() error {
	return nil
}

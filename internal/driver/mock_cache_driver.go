package driver

import (
	"context"
	"fmt"
	"time"

	driverErrors "mostafaqanbaryan.com/go-rest/internal/driver/errors"
)

type MockCacheDriver struct {
	list map[string]string
}

func NewMockCacheDriver() MockCacheDriver {
	return MockCacheDriver{
		list: make(map[string]string, 0),
	}
}

func (d MockCacheDriver) Get(ctx context.Context, key string) (string, error) {
	res, ok := d.list[key]
	if !ok {
		return "", driverErrors.ErrRecordNotFound
	}

	return res, nil
}

func (d MockCacheDriver) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	d.list[key] = fmt.Sprintf("%v", value)
	return nil
}

func (d MockCacheDriver) Close() error {
	return nil
}

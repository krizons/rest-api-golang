package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
	Close(ctx context.Context) error
}

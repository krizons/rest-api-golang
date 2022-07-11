package rediscache

import (
	"context"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestPutAndGet(t *testing.T) {
	assert := assert.New(t)
	cache, err := New("127.0.0.1:6379")
	assert.NoError(err)
	cache.Set("test", "val", 0)
	val, ok := cache.Get("test")
	assert.Equal(ok, true)
	if ok {
		data, err := redis.String(val, nil)
		assert.NoError(err)
		assert.Equal(data, "val")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	cache.Close(ctx)
}

package mycache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPutAndGet(t *testing.T) {
	assert := assert.New(t)
	cache, err := New(time.Second*1, time.Second*2)
	assert.NoError(err)
	cache.Set("test", "val", 0)
	val, ok := cache.Get("test")
	assert.Equal(ok, true)
	if ok {
		assert.Equal(val, "val")
	}
	time.Sleep(time.Second * 3)
	_, ok = cache.Get("test")
	assert.Equal(ok, false)
	cache.Close()
}

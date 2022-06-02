package rediscache

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}
type Cache struct {
	redis redis.Conn
}

func New(url string) (*Cache, error) {

	c, err := redis.Dial("tcp", url)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil, err
	}
	return &Cache{redis: c}, nil
}
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {

	_, err := c.redis.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

}
func (c *Cache) Get(key string) (interface{}, bool) {

	val, err := c.redis.Do("GET", key)

	if err != nil {
		return nil, false
	}
	if val == nil {
		return val, false
	}
	return val, true

}
func (c *Cache) Close() {
	c.redis.Close()
}

package redisc

import (
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/tsingson/bytecache/pkg/vtils"
)

type RedisConfig struct {
	Addr              string
	Password          string
	defaultExpiration time.Duration
	defaultDb         int
	debug             bool
}

var defaultRedisConfig = &RedisConfig{
	Password:          "vk_2018win",
	Addr:              "localhost:6379",
	defaultExpiration: 0,
	debug:             true,
}

func connectRedis(cfg *RedisConfig) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		// Password: cfg.Password,  // no password set
		DB: cfg.defaultDb, // use default DB
	})

	// _, err = client.Ping().Result()
	// fmt.Println(pong, err)
	// Output: PONG <nil>
	return
}

type RedisCache struct {
	redis             *redis.Client
	defaultExpiration time.Duration
}

func NewRedisCache(cfg *RedisConfig) (*RedisCache, error) {
	r := &RedisCache{defaultExpiration: cfg.defaultExpiration}

	client := connectRedis(cfg)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	r.redis = client
	return r, nil
}

func (c *RedisCache) Set(k, v []byte) bool {
	err := c.redis.Set(vtils.B2S(k), vtils.B2S(v), c.defaultExpiration).Err()
	check := true
	if err != nil {
		check = false
	}
	return check
}

func (c *RedisCache) Del(k []byte) {
	keys := []string{vtils.B2S(k)}

	err := c.redis.Del(keys...).Err()
	if err == nil {
		return
	}
}

func (c *RedisCache) Get(k []byte) ([]byte, bool) {
	var b []byte
	b, err := c.redis.Get(vtils.B2S(k)).Bytes()
	if err != nil {
		return nil, false
	}
	return b, true
}

// Clear clear
func (c *RedisCache) Clear() {
	_ = c.redis.FlushAll().Err()
}

// Save save
func (c *RedisCache) Save() {
}

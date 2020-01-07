package redisc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRedisCache(t *testing.T) {
	c, err := NewRedisCache(defaultRedisConfig)
	assert.NoError(t, err)

	k, v := []byte("key"), []byte("value")

	check := c.Set(k, v)

	assert.True(t, check)

	vv, cc := c.Get(k)
	assert.True(t, cc)
	if cc {
		assert.Equal(t, vv, v)
	}
}

/**
func BenchmarkRedisCache_Get(b *testing.B) {
	c, _ := NewRedisCache(defaultRedisConfig)

	k, v := []byte("key"), []byte("value")

	_ = c.Set(k, v)

	// b.SetParallelism(128)
	b.SetParallelism(runtime.NumCPU() * 2)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = c.Get(k)
		}
	})
}

*/

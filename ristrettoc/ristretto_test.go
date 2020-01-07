package ristrettoc

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tsingson/logger"
	"go.uber.org/zap/zapcore"
)

func TestNewGraphCache(t *testing.T) {
	var c *GraphCache
	var err error
	k, v := []byte("key"), []byte("value")

	t.Run("new", func(t *testing.T) {
		c, err = NewGraphCache(WithNumCounters(10000))
		assert.NoError(t, err)
	})
	t.Run("set", func(t *testing.T) {
		// t.Parallel()
		check := c.Set(k, v)
		assert.True(t, check)
	})
	t.Run("get", func(t *testing.T) {
		vv, check := c.Get(k)
		assert.True(t, check)
		assert.Equal(t, vv, v)
	})
	t.Run("del", func(t *testing.T) {
		c.Del(k)
		time.Sleep(time.Duration(10) * time.Microsecond)
		_, check := c.Get(k)
		assert.False(t, check)
	})
}

func BenchmarkGraphCache_Set(b *testing.B) {
	var c *GraphCache

	k, v := []byte("key"), []byte("value")
	c, _ = NewGraphCache()

	chk := c.Set(k, v)
	if chk {
		log := logger.New(logger.WithStoreInDay(),
			logger.WithDebug(),
			logger.WithDays(31),
			logger.WithPath("/home/go/bin/log"),
			logger.WithPrefix("ristretto"),
			logger.WithLevel(zapcore.DebugLevel))

		log.Error("log error chk")

		// b.SetParallelism(128)
		b.SetParallelism(runtime.NumCPU() * 4)
		b.StartTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, chk := c.Get(k)
				if !chk {
					log.Error("cache get not hit")
				}
			}
		})

		time.Sleep(time.Duration(5) * time.Second)
		log.Sync()
		log.Sync()

	}
}

/**
func TestGraphCache_Del2(t *testing.T) {
    fmt.Println("hello, and")
    fmt.Println("goodbye")
    // Output:
    // hello, and
    // goodbye
}

// 无序输出 Unordered output
func TestGraphCache_Del(t *testing.T) {
    for _, value := range []int{0,1,2,3} {
        fmt.Println(value)
    }
    // Unordered output: 4
    // 2
    // 1
    // 3
    // 0
}


*/

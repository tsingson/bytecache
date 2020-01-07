package fastc

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tsingson/logger"
	"go.uber.org/zap/zapcore"
)

/**
func TestMain(m *testing.M) {
    fmt.Println("Test Main")
    os.Exit(m.Serve())
}

*/

func TestNewFastCache(t *testing.T) {
	var c *VictoriaCache
	var err error
	k, v := []byte("key"), []byte("value1111")
	t.Run("new", func(t *testing.T) {
		c = NewVictoriaCache(WithCacheSize(1024 * 1024 * 512))
		assert.NoError(t, err)
	})
	t.Run("set ", func(t *testing.T) {
		check := c.Set(k, v)
		assert.True(t, check)
	})
	t.Run("get ", func(t *testing.T) {
		vv, check := c.Get(k)
		assert.True(t, check)
		if check {
			assert.Equal(t, vv, v)
		}
	})
	t.Run("del  ", func(t *testing.T) {
		c.Del(k)
		time.Sleep(time.Duration(10) * time.Microsecond)
		_, check := c.Get(k)
		assert.False(t, check)
	})
}

func BenchmarkVictoriaCache_Get(b *testing.B) {
	var c *VictoriaCache

	k, v := []byte("key"), []byte("value")
	c = NewVictoriaCache()

	chk := c.Set(k, v)

	if chk {

		log := logger.New(logger.WithStoreInDay(),
			logger.WithDebug(),
			logger.WithDays(31),
			logger.WithPath("/home/go/bin/log"),
			logger.WithPrefix("fastcache"),
			logger.WithLevel(zapcore.DebugLevel))

		// defer log.Sync()

		log.Error("log error chk")

		// b.SetParallelism(128)
		b.SetParallelism(runtime.NumCPU() * 2)
		b.StartTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, chk := c.Get(k)
				if !chk {
					log.Error("cache not hit")
				}
			}
		})
		time.Sleep(time.Duration(5) * time.Second)
		log.Sync()
		log.Sync()
	}
}

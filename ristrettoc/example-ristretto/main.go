package main

import (
	"os"
	"runtime"

	"github.com/tsingson/logger"
	"go.uber.org/zap/zapcore"

	"github.com/tsingson/bytecache/ristrettoc"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	log := logger.New(logger.WithStoreInDay(),
		logger.WithDebug(),
		logger.WithDays(31),
		logger.WithPath("/home/go/bin/log"),
		logger.WithPrefix("ristretto"),
		logger.WithLevel(zapcore.DebugLevel))

	var c *ristrettoc.GraphCache

	k, v := []byte("key"), []byte("value")
	c, _ = ristrettoc.NewGraphCache()

	chk := c.Set(k, v)
	log.Error("log error chk")

	if !chk {

		log.Error("log error chk")
		os.Exit(-1)
	}

	stopSignal := make(chan struct{})

	// defer log.Sync()

	get := func() {
		for i := 0; i < 1000; i++ {

			_, chk := c.Get(k)
			if !chk {
				log.Error("cache not hit")
			}
		}
	}
	go func() {
		get()
		log.Sync()
	}()
	go func() {
		get()
		log.Sync()
	}()
	go func() {
		get()
		log.Sync()
	}()

	<-stopSignal
}

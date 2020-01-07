package smallcache

import (
	"time"

	"github.com/VictoriaMetrics/fastcache"
)

const (
	defaultTimeOut   = time.Duration(10) * time.Microsecond
	defaultTTL       = time.Duration(24*30) * time.Hour
	defaultCacheSize = 1024 * 1024 * 512
)

// SmallCache fast cache from fast http
type SmallCache struct {
	cache       *fastcache.Cache
	cacheSize   int
	timeOut     time.Duration
	TTL         time.Duration
	StoragePath string
}

var defaultSmallCache = &SmallCache{
	cache:     nil,
	cacheSize: defaultCacheSize,
	timeOut:   defaultTimeOut,
	TTL:       defaultTTL,
}

// SmallCacheOption  options
type SmallCacheOption func(*SmallCache)

// WithCacheSize set cache size
func WithCacheSize(r int) SmallCacheOption {
	return func(o *SmallCache) {
		o.cacheSize = r
	}
}

func WithPath(r string) SmallCacheOption {
	return func(o *SmallCache) {
		o.StoragePath = r
	}
}

// WithCacheSize set cache size
func WithTTL(r time.Duration) SmallCacheOption {
	return func(o *SmallCache) {
		o.TTL = r
	}
}

// WithTimeOut set time out
func WithTimeOut(r time.Duration) SmallCacheOption {
	return func(o *SmallCache) {
		o.timeOut = r
	}
}

// NewSmallCache fast cache
func NewSmallCache(opts ...SmallCacheOption) *SmallCache {
	s := defaultSmallCache

	for _, o := range opts {
		o(s)
	}

	if len(s.StoragePath) > 0 {
		s.cache = fastcache.LoadFromFileOrNew(s.StoragePath, s.cacheSize)
	} else {
		s.cache = fastcache.New(s.cacheSize)
	}

	return s
}

// Set set
func (f *SmallCache) Set(k, v []byte) bool {
	f.cache.Set(k, v)
	time.Sleep(f.timeOut)
	return true
}

func (f *SmallCache) Del(k []byte) {
	f.cache.Del(k)
}

// Get get
func (f *SmallCache) Get(k []byte) (v []byte, check bool) {
	check = f.cache.Has(k)

	if check {
		v = f.cache.Get(nil, k)
	}
	return
}

// Clear clear
func (f *SmallCache) Clear() {
	f.cache.Reset()
}

// Save save
func (f *SmallCache) Save() error {
	if len(f.StoragePath) > 0 {
		return f.cache.SaveToFileConcurrent(f.StoragePath, 4)
	}
	return nil
}

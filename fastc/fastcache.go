package fastc

import (
	"runtime"
	"time"

	"github.com/VictoriaMetrics/fastcache"
)

// VictoriaCache fast cache from fast http
type VictoriaCache struct {
	path      string
	cache     *fastcache.Cache
	cacheSize int
	timeOut   time.Duration
	TTL       time.Duration
}

var defaultFastCache = &VictoriaCache{
	path:      defaultLogPath("fastcache"),
	cache:     nil,
	cacheSize: defaultCacheSize,
	timeOut:   defaultTimeOut,
	TTL:       defaultTTL,
}

// VictoriaCacheOption  options
type VictoriaCacheOption func(*VictoriaCache)

// WithCacheSize set cache size
func WithCacheSize(r int) VictoriaCacheOption {
	return func(o *VictoriaCache) {
		o.cacheSize = r
	}
}

func WithPath(p string) VictoriaCacheOption {
	return func(o *VictoriaCache) {
		o.path = defaultLogPath("fastcache") + "/" + p
	}
}

// WithCacheSize set cache size
func WithTTL(r time.Duration) VictoriaCacheOption {
	return func(o *VictoriaCache) {
		o.TTL = r
	}
}

// WithTimeOut set time out
func WithTimeOut(r time.Duration) VictoriaCacheOption {
	return func(o *VictoriaCache) {
		o.timeOut = r
	}
}

// NewVictoriaCache fast cache
func NewVictoriaCache(opts ...VictoriaCacheOption) *VictoriaCache {
	s := defaultFastCache

	for _, o := range opts {
		o(s)
	}

	if len(s.path) > 0 {
		s.cache = fastcache.LoadFromFileOrNew(s.path, s.cacheSize)
	} else {
		s.cache = fastcache.New(s.cacheSize)
	}

	return s
}

// Set set
func (c *VictoriaCache) Set(k, v []byte) bool {
	c.cache.SetBig(k, v)
	time.Sleep(c.timeOut)
	return true
}

// Del delete
func (c *VictoriaCache) Del(k []byte) {
	c.cache.Del(k)
}

// Get get
func (c *VictoriaCache) Get(k []byte) (v []byte, check bool) {
	v, check = c.cache.HasGet(nil, k)

	if check {
		v = c.cache.GetBig(nil, k)
	}
	return
}

// Clear clear
func (c *VictoriaCache) Clear() {
	c.cache.Reset()
}

// Save save
func (c *VictoriaCache) Save() {
	_ = c.cache.SaveToFileConcurrent(c.path, runtime.NumCPU())
}

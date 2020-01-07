package ristrettoc

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

const valueCost = 8

var defaultGraphCacheConfdig = ristretto.Config{
	NumCounters: 1e8,         // 1e7 number of keys to track frequency of (10M).
	MaxCost:     2 * 1 << 30, // maximum cost of cache (1GB).
	BufferItems: 64 * 4,      // number of keys per Get buffer.
	Metrics:     false,
	// Hashes:      1,
}

// GraphCacheOption options
type GraphCacheOption func(*ristretto.Config)

// WithNumCounters options
func WithNumCounters(r int64) GraphCacheOption {
	return func(o *ristretto.Config) {
		o.NumCounters = r
	}
}

// WithMaxCost options
func WithMaxCost(m int64) GraphCacheOption {
	return func(o *ristretto.Config) {
		o.MaxCost = m
	}
}

// WithBufferItems options
func WithBufferItems(m int64) GraphCacheOption {
	return func(o *ristretto.Config) {
		o.BufferItems = m
	}
}

// GraphCache DGraphCache cache
type GraphCache struct {
	cache     *ristretto.Cache
	timeOut   time.Duration
	valueCost int64
}

// NewGraphCache new
func NewGraphCache(opts ...GraphCacheOption) (*GraphCache, error) {
	cfg := &defaultGraphCacheConfdig
	for _, o := range opts {
		o(cfg)
	}

	c, err := ristretto.NewCache(cfg)
	if err != nil {
		return nil, err
	}
	return &GraphCache{cache: c, timeOut: defaultTimeOut}, nil
}

// Set set
func (c *GraphCache) Set(k, v []byte) bool {
	ck := c.cache.Set(k, v, valueCost)
	// wait for value to pass through buffers
	time.Sleep(c.timeOut)
	return ck
}

// Del del cache via key
func (c *GraphCache) Del(k []byte) {
	c.cache.Del(k)
}

// Get get
func (c *GraphCache) Get(k []byte) (v []byte, ck bool) {
	var vv interface{}
	vv, ck = c.cache.Get(k)
	if ck {
		v = vv.([]byte)
	}
	return v, ck
}

// Clear clear
func (c *GraphCache) Clear() {
	c.cache.Clear()
}

// Save save
func (c *GraphCache) Save() {
}

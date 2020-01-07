package ristrettoc

import (
	"time"
)

const (
	wait             = time.Duration(10) * time.Millisecond
	defaultCacheSize = 1024 * 1024 * 1024
	defaultTimeOut   = time.Duration(100) * time.Nanosecond
	defaultTTL       = time.Duration(24*30) * time.Hour
)

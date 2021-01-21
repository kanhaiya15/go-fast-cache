package store

import (
	"time"

	"github.com/kanhaiya15/go-fast-cache/cfg"
	"github.com/patrickmn/go-cache"
)

// Setup new cache
func Setup() *cache.Cache {
	return cache.New(time.Duration(cfg.CacheInMinute)*time.Minute, time.Duration(cfg.CachePurgeInMinute)*time.Minute)
}

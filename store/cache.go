package store

import (
	"fmt"
	"os"
	"strconv"
	"time"

	cache "github.com/patrickmn/go-cache"
)

// Setup new cache
func Setup() *cache.Cache {
	cacheInMinute := 10
	cachePurgeInMinute := 5

	cacheInMinute, err := strconv.Atoi(os.Getenv("CACHE_IN_MINUTE"))
	if err != nil {
		fmt.Println("error in getting cacheInMinute from env, setting up default ")
	}
	cachePurgeInMinute, err = strconv.Atoi(os.Getenv("CACHE_PURGE_IN_MINUTE"))
	if err != nil {
		fmt.Println("error in getting cachePurgeInMinute from env, setting up default ")
	}
	return cache.New(time.Duration(cacheInMinute)*time.Minute, time.Duration(cachePurgeInMinute)*time.Minute)
}

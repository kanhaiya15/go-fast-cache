package cfg

import (
	"fmt"
	"os"
	"strconv"
)

// Default
var (
	ServerPort         = 8080
	CacheInMinute      = 10
	CachePurgeInMinute = 5
)

// Setup cfg
func Setup() {
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		fmt.Println("error in getting ServerPort from env, setting up default : ", ServerPort)
	} else {
		ServerPort = serverPort
	}

	cacheInMinute, err := strconv.Atoi(os.Getenv("CACHE_IN_MINUTE"))
	if err != nil {
		fmt.Println("error in getting CacheInMinute from env, setting up default : ", CacheInMinute)
	} else {
		CacheInMinute = cacheInMinute
	}
	cachePurgeInMinute, err := strconv.Atoi(os.Getenv("CACHE_PURGE_IN_MINUTE"))
	if err != nil {
		fmt.Println("error in getting CachePurgeInMinute from env, setting up default : ", CachePurgeInMinute)
	} else {
		CachePurgeInMinute = cachePurgeInMinute
	}
}

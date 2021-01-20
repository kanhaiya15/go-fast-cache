package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/go-fast-cache/store"
	"github.com/kanhaiya15/go-fast-cache/types"
	"github.com/patrickmn/go-cache"
)

var (
	newCache *cache.Cache
)

func init() {
	newCache = store.Setup()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GoLang", "status": http.StatusOK})
	})

	router.POST("/post", func(c *gin.Context) {
		var json types.Entity
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newCache.Set(fmt.Sprint(json.ID), &json, cache.DefaultExpiration)
		c.JSON(http.StatusOK, gin.H{"message": "OK", "status": http.StatusOK})
	})

	router.GET("/post/:id", func(c *gin.Context) {
		if x, ok := newCache.Get(fmt.Sprint(c.Param("id"))); ok {
			entity := x.(*types.Entity)
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"status":  http.StatusOK,
				"results": map[string]interface{}{
					"id":   entity.ID,
					"name": entity.Name,
				}})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "0 results found", "status": http.StatusOK})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

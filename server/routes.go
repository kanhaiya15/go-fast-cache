package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/go-fast-cache/cfg"
	"github.com/kanhaiya15/go-fast-cache/models"
	"github.com/kanhaiya15/go-fast-cache/store"
	"github.com/patrickmn/go-cache"
)

var (
	newCache *cache.Cache
)

func init() {
	cfg.Setup()
	newCache = store.Setup()
}

// Setup routes
func Setup() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GoLang", "status": http.StatusOK})
	})

	router.POST("/post", func(c *gin.Context) {
		var json models.Post
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newCache.Set(fmt.Sprint(json.ID), &json, cache.DefaultExpiration)
		c.JSON(http.StatusOK, gin.H{"message": "OK", "status": http.StatusOK})
	})

	router.GET("/post/:id", func(c *gin.Context) {
		if x, ok := newCache.Get(fmt.Sprint(c.Param("id"))); ok {
			entity := x.(*models.Post)
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
	return router
}

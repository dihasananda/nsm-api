package router

import (
	"net/http"
	"nsm-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(schoolHandler *handlers.SchoolHandler) *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/schools", schoolHandler.GetSchools)

	// NoRoute handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "The requested resource does not exist",
		})
	})

	return r
}

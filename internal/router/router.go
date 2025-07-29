package router

import (
	"nsm-api/configs"
	"nsm-api/internal/handlers"
	"nsm-api/internal/repository"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// DB connection from configs
	db := configs.InitDB()

	// Repository & Handler
	NSMRepo := repository.NewNSMRepository(db)
	NSMHandler := handlers.NewNsmHandler(NSMRepo)

	// Routes
	r.GET("/health", handlers.HealthCheck)
	r.GET("/ping", handlers.Ping)
	r.GET("/nsms", NSMHandler.GetNSMs)

	// NoRoute handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "The requested resource does not exist",
		})
	})

	return r
}

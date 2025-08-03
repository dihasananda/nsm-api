package main

import (
	"nsm-api/configs"
	"nsm-api/internal/handlers"
	"nsm-api/internal/repository"
	"nsm-api/internal/router"
	"nsm-api/internal/services"
)

func main() {
	// Init DB
	db := configs.InitDB()

	// Dependencies
	schoolRepo := repository.NewSchoolRepository(db)
	schoolService := services.NewSchoolService(schoolRepo)
	schoolHandler := handlers.NewSchoolHandler(schoolService)

	// Router
	r := router.SetupRouter(schoolHandler)

	// Run server
	r.Run(":8080")
}

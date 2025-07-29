package main

import (
	"log"
	"nsm-api/internal/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

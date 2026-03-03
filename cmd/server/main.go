package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/sakshya/gateway-tee/internal/api"
)

func main() {
	log.Println("Initializing Sakshya TEE Gateway...")

	r := gin.Default()

	// Apply security middleware
	r.Use(api.CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/verify", api.VerifyHandler)
	}

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

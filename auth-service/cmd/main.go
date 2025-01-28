package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/morshedulmunna/auth-service/internal/handler"
)

func main() {
	r := gin.Default()

	// Initialize handlers
	authHandler := handler.NewAuthHandler()

	// Routes
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	log.Fatal(r.Run(":8080"))
}

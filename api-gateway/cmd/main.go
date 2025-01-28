package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morshedulmunna/api-gateway/internal/config"
	"github.com/morshedulmunna/api-gateway/internal/handler"
	"github.com/morshedulmunna/api-gateway/internal/middleware"
)

func main() {
	r := gin.Default()

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Gateway Service",
		})
	})

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Load configuration
	config, err := config.LoadConfig("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize handlers
	gatewayHandler := handler.NewGatewayHandler(config)

	// Public routes (no auth required)
	auth := r.Group("/api/auth")
	{
		for _, route := range config.Services["auth"].Routes {
			auth.Handle(route.Method, route.Path, gatewayHandler.ProxyAuthService)
		}
	}

	// Protected routes (auth required)
	users := r.Group("/api/users")
	users.Use(middleware.AuthMiddleware())
	{
		for _, route := range config.Services["users"].Routes {
			users.Handle(route.Method, route.Path, gatewayHandler.ProxyUserService)
		}
	}

	log.Fatal(r.Run(":8000"))
}

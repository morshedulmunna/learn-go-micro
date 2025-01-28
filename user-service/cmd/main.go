package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/morshedulmunna/user-service/internal/handler"
)

func main() {
	r := gin.Default()

	// Initialize handlers
	userHandler := handler.NewUserHandler()

	// Routes
	users := r.Group("/users")
	{
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.GET("/", userHandler.ListUsers)
	}

	log.Fatal(r.Run(":8081"))
}

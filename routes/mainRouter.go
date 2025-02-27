package routes

import (
	"fmt"
	"golang-rnd/handlers"
	"golang-rnd/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)

	// Protected route with JWT verification
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/data", handlers.DataHandler)
	}

	fmt.Println("Server is running on :8080")
	r.Run(":8080")
}

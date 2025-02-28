package routes

import (
	"fmt"
	"golang-rnd/handlers"
	"golang-rnd/initializers"
	"golang-rnd/middleware"
	"golang-rnd/schema"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router() {
	config := initializers.LoadConfig()
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)
	r.POST("/validate",
		middleware.ValidateBody[schema.ILoginReq](),
		middleware.ValidateProxy(),
		middleware.ValidateMatrix(middleware.Read),
		handlers.SampleController)

	// Protected route with JWT verification
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/data", handlers.DataHandler)
	}

	port := strconv.Itoa(config.Port)
	fmt.Println("✅ Server is running on port:", port)

	r.Run(":" + port)
}

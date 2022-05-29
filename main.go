package main

import (
	"os"

	"github.com/RAGAZZ-Studio/liquiz-backend/app/controllers"
	"github.com/RAGAZZ-Studio/liquiz-backend/app/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

  r.Use(cors.Default())
	r.Run(":" + port)
}

package main

import (
	"github.com/RAGAZZ-Studio/liquiz-backend/app/controllers"
	"github.com/RAGAZZ-Studio/liquiz-backend/app/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}

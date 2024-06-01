package main

import (
	"myapp/config"
	"myapp/controllers"
	"myapp/middleware"
	"myapp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to the database
	config.ConnectDB()

	// Migrate the schema
	var user models.User
	var photo models.Photo
	user.Migrate(config.DB)
	photo.Migrate(config.DB)

	// Routes
	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/photos", controllers.AddPhoto)
		authorized.DELETE("/photos", controllers.DeletePhoto)
	}

	// Start server
	r.Run(":8080")
}

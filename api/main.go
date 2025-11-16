package main

import (
	"go_crud/api/controllers"
	"go_crud/api/initializers"
	"go_crud/api/modules"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	v1 := r.Group("/api/v1")

	modules.RegisterUserRoutes(v1)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go CRUD API",
		})

	})

	// Define a simple GET endpoint
	v1.POST("/posts", controllers.Postscreate)
	v1.GET("/posts", controllers.PostsIndex)
	v1.GET("/posts/:id", controllers.PostsShow)
	v1.DELETE("/posts/:id", controllers.PostsDelete)
	v1.PUT("/posts/:id", controllers.PostsUpdate)

	r.Run()
}

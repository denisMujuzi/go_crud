package main

import (
	"go_crud/api/controllers"
	"go_crud/api/initializers"
	"go_crud/api/modules"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Redis_connection()
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
	r.GET("/redis-check/:id/:value", func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")
		value := c.Param("value")

		val, err := initializers.Rdb.Get(ctx, id).Result()

		if err != nil {
			err := initializers.Rdb.Set(ctx, id, value, 5*time.Second).Err()
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, gin.H{"msg": "field set"})
			return
		}
		c.JSON(200, gin.H{id: val})

	})

	// Define a simple GET endpoint
	v1.POST("/posts", controllers.Postscreate)
	v1.GET("/posts", controllers.PostsIndex)
	v1.GET("/posts/:id", controllers.PostsShow)
	v1.DELETE("/posts/:id", controllers.PostsDelete)
	v1.PUT("/posts/:id", controllers.PostsUpdate)

	r.Run()
}

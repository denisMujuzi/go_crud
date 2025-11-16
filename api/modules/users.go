package modules

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	r := router.Group("/users")
	users_data := map[string]string{
		"1": "Alice",
		"2": "Bob",
		"3": "Charlie",
	}

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		if name, exists := users_data[id]; exists {
			c.JSON(200, gin.H{
				"user_id":   id,
				"user_name": name,
			})
			return
		}

		c.JSON(404, gin.H{
			"error": fmt.Sprintf("User with ID %s not found", id),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, users_data)

	})

	r.PUT("/:id", func(ctx *gin.Context) {
		type BodyType struct {
			Name *string
			Id   *string
		}

		var Body BodyType

		ctx.Bind(&Body)
		if Body.Name == nil || Body.Id == nil {
			ctx.JSON(400, gin.H{"error": "Name and Id are required"})
			return
		}

		// check if id exists
		if _, exists := users_data[*Body.Id]; !exists {
			ctx.JSON(404, gin.H{
				"error": fmt.Sprintf("User with ID %s not found", *Body.Id),
			})
			return
		}

		// update user

		users_data[*Body.Id] = *Body.Name
		ctx.JSON(200, gin.H{
			"message": "User updated successfully",
			"user":    users_data,
		})
	})
}

package main

import (
	"go_crud/api/controllers"
	"go_crud/api/initializers"
	"go_crud/api/modules"

	"net/http"

	"github.com/gin-gonic/gin"
	bridge "github.com/vercel/go-bridge/go/bridge"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// This is required by Vercel, do not rename it
func __NOW_HANDLER_FUNC_NAME(w http.ResponseWriter, r *http.Request) {
	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery())

	v1 := g.Group("/api/v1")
	modules.RegisterUserRoutes(v1)

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go CRUD API",
		})
	})

	v1.POST("/posts", controllers.Postscreate)
	v1.GET("/posts", controllers.PostsIndex)
	v1.GET("/posts/:id", controllers.PostsShow)
	v1.DELETE("/posts/:id", controllers.PostsDelete)
	v1.PUT("/posts/:id", controllers.PostsUpdate)

	g.ServeHTTP(w, r)
}

func main() {
	bridge.Start(http.HandlerFunc(__NOW_HANDLER_FUNC_NAME))
}

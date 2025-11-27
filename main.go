package main

import (
	"go_crud/api/controllers"
	"go_crud/api/initializers"
	"go_crud/api/modules"
	"go_crud/rabbitmq"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Redis_connection()
	// initializers.RabbitMQ_connection()
	initializers.InitializeFirebaseApp()
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

	// send message to rabbitmq
	r.POST("/rabbitmq/:msg", func(ctx *gin.Context) {
		msg := ctx.Param("msg")

		ch := initializers.RabbitMQChannel

		err := ch.PublishWithContext(ctx,
			"",           // exchange
			"task_queue", // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(msg),
			})
		initializers.FailOnError(err, "Failed to publish a message")
		ctx.JSON(200, gin.H{"message": "Message sent to RabbitMQ"})
	})

	r.POST("/u-ex/:log", func(ctx *gin.Context) {
		log := ctx.Param("log")

		ch := initializers.RabbitMQChannel

		err := ch.PublishWithContext(ctx,
			"logs", // exchange
			"",     // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(log),
			})
		initializers.FailOnError(err, "Failed to publish a message")
		ctx.JSON(200, gin.H{"message": "Message sent to RabbitMQ Exchange"})
	})

	r.GET("/firebase/:msg", func(ctx *gin.Context) {
		msg := ctx.Param("msg")
		initializers.SendMessage(msg)
		ctx.JSON(200, gin.H{"message": "Message sent to Firebase"})

	})

	go rabbitmq.StartConsumers()

	// Define a simple GET endpoint
	v1.POST("/posts", controllers.Postscreate)
	v1.GET("/posts", controllers.PostsIndex)
	v1.GET("/posts/:id", controllers.PostsShow)
	v1.DELETE("/posts/:id", controllers.PostsDelete)
	v1.PUT("/posts/:id", controllers.PostsUpdate)

	r.Run()
}

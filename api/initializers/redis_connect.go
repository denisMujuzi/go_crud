package initializers

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Redis_connection() {
	ctx := context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// check for error
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		// Optionally log error here
		return
	}

}

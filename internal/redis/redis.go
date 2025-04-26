package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func ConnectRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Use the Redis container or your Redis service endpoint
	})
}

func CheckUsername(username string) bool {
	ctx := context.Background()
	exists, err := client.Exists(ctx, username).Result()
	if err != nil {
		log.Println("Redis error:", err)
		return false
	}
	return exists > 0
}

func SetUsername(username string) {
	ctx := context.Background()
	err := client.Set(ctx, username, "exists", 24*time.Hour).Err() // Cache for 24 hours
	if err != nil {
		log.Println("Failed to set username in Redis:", err)
	}
}

package initializers

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Redis 		 *redis.Client
	RedisContext context.Context
)

func ConnectToRedis(){
	RedisContext = context.Background()

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
        Password: os.Getenv("REDIS_PASSWORD"),
        DB:       0,
	})
}
package database

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

func NewRedis() *Cache {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Panic(err)
	}
	return &Cache{
		redis.NewClient(
			&redis.Options{
				Addr:     os.Getenv("REDIS_ADDR"),
				Password: os.Getenv("REDIS_PASS"),
				DB:       redisDB,
			},
		),
	}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}

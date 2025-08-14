package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)


var (
	Rdb *redis.Client
) 

func ConnectRedis() {
	_ = godotenv.Load()
	addr := os.Getenv("REDIS_ADDRESS")
	Password := os.Getenv("REDIS_PASSWORD")

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: Password, 
		DB:       0,        
	})
}

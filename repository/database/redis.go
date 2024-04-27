package database

import (
	"github.com/redis/go-redis/v9"
	"storage/config"
)

func ConnectToRedis() *redis.Client {
	cfg := config.GetRedis()
	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Addrs[0], //"localhost:6379",
		Password: cfg.Password, // no password set
		DB:       0,            // use default DB
	})
	return db
}

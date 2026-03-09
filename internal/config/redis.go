package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redis'e Baglanilamadi: %v", err)
	}

	log.Println("Redis Baglantisi kuruldu")

	return client
}

package server

import (
	"catalog/src/config"
	"github.com/go-redis/redis/v7"
	"time"
)


func NewRedisCache(config *config.Config) *redis.Client {

	return redis.NewClient(&redis.Options{
		Password:     config.Server.RedisPassword,
		Addr:         config.Server.RedisUrl,
		MinIdleConns: config.Server.RedisMaxIdle,
		IdleTimeout:  time.Duration(config.Server.RedisIdleTimeout) * time.Second,
		DB:           config.Server.RedisDatabase,
	})

}


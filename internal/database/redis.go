package database

import (
	"github.com/go-redis/redis"
)

type RedisClient struct {
	Client *redis.Client
}

type RedisOptions struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(opts *RedisOptions) *RedisClient {
	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     opts.Addr,
			Password: opts.Password,
			DB:       opts.DB,
		}),
	}
}

func (c *RedisClient) CloseRedisClient() {
	c.Client.Close()
}

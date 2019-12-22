package db

import (
	"github.com/go-redis/redis/v7"

	"github.com/anvari1313/yaus/config"
)

func ConnectRedis(c config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:               c.Address,
		Password:           c.Password,
		DB:                 c.DB,
		MaxRetries:         c.MaxRetries,
		MinRetryBackoff:    c.MinRetryBackOff,
		MaxRetryBackoff:    c.MaxRetryBackOff,
		DialTimeout:        c.DialTimeout,
		ReadTimeout:        c.ReadTimeout,
		WriteTimeout:       c.WriteTimeout,
		PoolSize:           c.PoolSize,
		MinIdleConns:       c.MinIdleConnections,
		MaxConnAge:         c.MaxConnectionAge,
		PoolTimeout:        c.PoolTimeout,
		IdleTimeout:        c.IdleTimeout,
		IdleCheckFrequency: c.IdleCheckFrequency,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

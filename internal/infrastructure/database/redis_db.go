package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(addr, psswd string, ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: psswd,
		DB:       0,
		Protocol: 2,
	})

	val, err := client.Get(ctx, "foo").Result()
	if err != redis.Nil {
		log.Fatal(err)
	}
	fmt.Println("foo", val)

	return client
}

package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(addr, pwd string, ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
		Protocol: 2,
	})

	val, err := client.Get(ctx, "foo").Result()
	if !errors.Is(err, redis.Nil) {
		log.Fatal(err)
	}
	fmt.Println("foo", val)

	return client
}

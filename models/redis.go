package models

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var ctx = context.Background()

func init() {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Rdb.Ping(ctx)
}

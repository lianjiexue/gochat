package models

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Rdbctx = context.Background()

func init() {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	res := Rdb.Ping(Rdbctx)
	log.Println(res.Result())
}

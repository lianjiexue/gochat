package modeltest

import (
	"app/models"
	"context"
	"log"
	"testing"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {

	log.Println(ctx)
	t.Log("执行成功")
}
func TestRdb(t *testing.T) {
	models.Rdb.Set(ctx, "hello", "guo", 12)
	t.Log(models.Rdb.Get(ctx, "hello"))
}
func TestHello(t *testing.T) {
	t.Log("hello")
}

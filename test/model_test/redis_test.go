package modeltest

import (
	"app/models"
	"log"
	"testing"
)

func TestPing(t *testing.T) {

	log.Println(models.Rdb.Ping(models.Rdbctx))
	t.Log("执行成功")
}
func TestRdbSet(t *testing.T) {
	models.Rdb.Set(models.Rdbctx, "hello", "guo", 12)
	t.Log(models.Rdb.Get(models.Rdbctx, "hello"))
}
func TestLpuss(t *testing.T) {
	models.Rdb.LPush(models.Rdbctx, "messages", 123)
	t.Log("ok")
}
func TestSaveMsg(t *testing.T) {
	models.SaveMessage()
	t.Log("ok")
}

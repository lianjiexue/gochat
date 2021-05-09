package model_test

import (
	"app/model"
	"testing"
)

func TestList(t *testing.T) {
	lists := model.GetFUllMoods(1, 5)
	t.Log(lists)
	t.Log("ok")
}

package model_test

import (
	"app/models"
	"testing"
)

func TestList(t *testing.T) {
	lists := models.GetFUllMoods(1, 5)
	t.Log(lists)
	t.Log("ok")
}

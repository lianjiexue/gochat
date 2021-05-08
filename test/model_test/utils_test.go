package modeltest

import (
	"app/utils"
	"testing"
)

func TestGetDateString(t *testing.T) {
	t.Log(utils.GetDateString())
	t.Log("执行成功")
}

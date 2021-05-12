package project_test

import (
	"app/lib"
	"testing"
)

func TestCreateToken(t *testing.T) {
	//测试生成token
	tokenStr := lib.CreateToken(12)
	//输出token
	t.Log(tokenStr)
	//验证token
	t.Log(lib.ValidToken(tokenStr))
}

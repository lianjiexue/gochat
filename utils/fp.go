package utils

import (
	"math/rand"
	"time"
)

//获取日期的字符串表示
func GetDateString() string {
	date := time.Now().Format("20060102150405")
	return date
}

//返回随机数
func GetRandom(long int) int {
	rand.Seed(time.Now().Unix())
	return int(rand.Intn(1000))
}

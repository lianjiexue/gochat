package model

import (
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var db *gorm.DB

func init(){
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/zhuanyun?charset=utf8mb4&parseTime=True&loc=Local"
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
    	log.Println("数据库连接失败")
    	panic(err)
    }
}
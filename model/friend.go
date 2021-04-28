package model

import (
	"log"
)

type Friend struct {
	Uid int    `gorm:"uid" json:"uid"`
	Fid string `gorm:"fid" json:"fid"`
}

func (u *Friend) TableName() string {
	return "gc_friends"
}

func getFriendsById(uid int) []User {
	var fids []int
	db.Table("gc_friends").Where("uid", uid).Pluck("fid", &fids)
	log.Println(fids)
	var users []User
	db.Where("id IN ?", fids).Find(&users)
	return users
}

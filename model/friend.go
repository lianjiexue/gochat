package model

import (
	"log"
)

type Friend struct {
	Uid int `gorm:"uid" json:"uid"`
	Fid int `gorm:"fid" json:"fid"`
}

func (u *Friend) TableName() string {
	return "gc_friends"
}
func GetFriendsById(uid int) []User {
	var fids []int
	db.Table("gc_friends").Where("uid", uid).Pluck("fid", &fids)
	log.Println(fids)
	var users []User
	db.Where("id IN ?", fids).Find(&users)
	return users
}
func getFriendsById(uid int) []User {
	var fids []int
	db.Table("gc_friends").Where("uid", uid).Pluck("fid", &fids)
	log.Println(fids)
	var users []User
	db.Where("id IN ?", fids).Find(&users)
	return users
}

//移除好友
func RemoveFriend(uid int, fid int) bool {

	affected := db.Where("uid", uid).Where("fid", fid).Delete(&Friend{})
	if affected.Error == nil {
		return true
	}
	return true
}

//添加好友

func AddFriend(uid int, fid int) bool {
	// 执行关注
	var newFriend Friend

	newFriend.Uid = uid
	newFriend.Fid = fid

	affrected := db.Model(&Friend{}).Create(&newFriend)
	if affrected.Error == nil {
		return true
	}
	return false
}

//判断是否已经是好友了

func IsFriend(uid int, fid int) bool {
	var friend Friend
	result := db.Model(&Friend{Uid: uid, Fid: fid}).Take(&friend)
	if result.Error == nil {
		return false
	}
	return true
}

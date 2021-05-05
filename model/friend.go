package model

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Friend struct {
	Uid int `gorm:"uid" json:"uid"`
	Fid int `gorm:"fid" json:"fid"`
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
func UnFollow(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	fid := ctx.PostForm("fid")
	newUid, _ := strconv.Atoi(uid)
	newFid, _ := strconv.Atoi(fid)
	var friend Friend
	result := db.Where("uid", newUid).Where("fid", newFid).Take(&friend)
	log.Println(result)
	log.Print(friend)
	log.Println(newFid, newUid)
	if result.Error != nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "未关注",
		})
		return
	}
	friend.Uid = newUid
	friend.Fid = newFid
	affected := db.Where("uid", newUid).Where("fid", newFid).Delete(&friend)

	log.Println(affected)
	if affected.Error != nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "error",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
	}
}

func Follow(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	fid := ctx.PostForm("fid")
	newUid, _ := strconv.Atoi(uid)
	newFid, _ := strconv.Atoi(fid)
	// 执行查询 看是否已经关注
	var friend Friend
	result := db.Model(&Friend{Uid: newUid, Fid: newFid}).Take(&friend)
	if result.Error != nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "error",
		})
		return
	}
	// 执行关注
	var newFriend Friend

	newFriend.Uid = newUid
	newFriend.Fid = newFid

	afreact := db.Model(&Friend{}).Create(&newFriend)
	if afreact.Error != nil {
		ctx.JSON(200, gin.H{
			"code": 0,
		})

	} else {

		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
	}

}

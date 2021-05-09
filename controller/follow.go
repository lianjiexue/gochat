package controller

import (
	"app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//我的关注的用户
func MyFollow(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, error := strconv.Atoi(uid)
	if error != nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "error",
		})
	}
	friends := model.GetFriendsById(newUid)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    friends,
	})
}

//关注
func Follow(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	fid := ctx.PostForm("fid")
	newUid, _ := strconv.Atoi(uid)
	newFid, _ := strconv.Atoi(fid)
	// 执行查询 看是否已经关注
	is_friend := model.IsFriend(newUid, newFid)
	if is_friend {
		ctx.JSON(200, gin.H{
			"code": 200,
		})
		return
	}
	//添加好友
	is_ok := model.AddFriend(newUid, newFid)

	if is_ok {
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

//取消关注
func Unfollow(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	fid := ctx.PostForm("fid")
	newUid, _ := strconv.Atoi(uid)
	newFid, _ := strconv.Atoi(fid)
	//
	is_remove := model.RemoveFriend(newUid, newFid)

	if is_remove {
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

package controllers

import (
	"app/models"
	"app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//用户信息
func UserInfo(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	user := models.GetUserInfo(newUid)
	if user.Id != 0 {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    user,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "success",
		})
	}
}

//我的好友
func MyFriends(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	users := models.GetFriendsById(newUid)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    users,
	})
}

//查看当前用户返回用户是否被关注

func UserFollow(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	fid := ctx.PostForm("fid")
	newFid, _ := strconv.Atoi(fid)

	user := models.GetUserInfo(newUid)
	is_follow := models.IsFriend(newUid, newFid)

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"user":      user,
			"is_follow": is_follow,
		},
	})

}

//更新昵称
func UpdateNickname(ctx *gin.Context) {
	nickname := ctx.PostForm("nickname")
	uid := ctx.PostForm("uid")
	newuid, _ := strconv.Atoi(uid)
	is_update := models.SetUserNickname(newuid, nickname)
	if is_update {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "fail",
		})
	}
}

//更新头像
func UpdateHeadImg(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	uid := ctx.PostForm("uid")
	newUid, error := strconv.Atoi(uid)
	if error != nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "error",
		})
		return
	}
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "error",
		})
		return
	}
	fileRes := utils.TouchFilePath(file)
	if fileRes.Code == 0 {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "保存头像失败",
		})
		return
	}
	ctx.SaveUploadedFile(file, fileRes.Dst)
	dm := "https://chat.daguozhensi.com"
	headImg := dm + fileRes.SavePath
	is_ok := models.UpdateHeadImg(newUid, headImg)
	if is_ok {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"head_img": headImg,
			},
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "fail",
		})
	}
}

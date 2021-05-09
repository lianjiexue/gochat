package controller

import (
	"app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获取一条心情记录
func OneMood(ctx *gin.Context) {
	mood := model.OneMood()

	if mood.Id != 0 {
		user := model.GetUserInfo(mood.Uid)

		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "succcess",
			"data": gin.H{
				"user": user,
				"mood": mood,
			},
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "succcess",
		})
	}

}

//添加记录
func AddMood(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	content := ctx.PostForm("content")
	is_write := model.AddMood(newUid, content)
	if is_write {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

//删除记录
func DelMood(ctx *gin.Context) {
	id := ctx.PostForm("id")
	newId, _ := strconv.Atoi(id)
	is_del := model.DelMood(newId)
	if is_del {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

package controllers

import (
	"app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//我发的所有心情
func MoodsLists(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	page := ctx.PostForm("page")
	newUid, _ := strconv.Atoi(uid)
	newPage, _ := strconv.Atoi(page)
	moods := models.GetFUllMoods(newUid, newPage)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    moods,
	})
}

//获取一条心情记录
func OneMood(ctx *gin.Context) {
	mood := models.OneMood()

	if mood.Id != 0 {
		user := models.GetUserInfo(mood.Uid)

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
	is_write := models.AddMood(newUid, content)
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
	is_del := models.DelMood(newId)
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

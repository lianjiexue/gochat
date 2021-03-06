package controllers

import (
	"app/lib"
	"app/models"

	"github.com/gin-gonic/gin"
)

//登录
func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	user := models.GetOneUser(email, password)
	if user.Id != 0 {
		friends := models.GetFriendsById(user.Id)
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"user":    user,
				"friends": friends,
				"token":   lib.CreateToken(user.Id),
			},
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "fail",
		})
	}
}

//注册
func Register(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	verify := ctx.PostForm("verify")
	sex := ctx.PostForm("sex")
	if verify != "千里江陵一日还" {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "验证码错误",
		})
	}
	user := models.AddUser(email, password, sex)
	if user.Id != 0 {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    user,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "fail",
		})
	}
}

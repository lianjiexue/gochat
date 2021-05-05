package model

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `gorm:"uid" json:"id"`
	Nickname string `gorm:"nickname" json:"nickname"`

	HeadImg     string   `gorm:"head_img" json:"head_img"`
	Description string   `gorm:"description" json:"description"`
	Friends     []Friend `gorm:"foreignKey:uid" json:"firends,omitempty"`
}
type OneUser struct {
	Id       int    `gorm:"uid" json:"id"`
	Nickname string `gorm:"nickname" json:"nickname"`
	HeadImg  string `gorm:"head_img" json:"head_img"`
}
type RegisterUser struct {
	Id       string `gorm:"id" json:"id"`
	Nickname string `gorm:"nickname" json:"nickname"`
	Password string `gorm:"password" json:"password"`
	Email    string `gorm:"email" json:"email"`
	HeadImg  string `gorm:"head_img" json:"head_img"`
}

func (u *User) TableName() string {
	return "gc_users"
}
func (ou *OneUser) TableName() string {
	return "gc_users"
}
func (ou *RegisterUser) TableName() string {
	return "gc_users"
}
func GetUserName(uid int) string {
	var user User
	db.Where("id=?", uid).First(&user)
	return user.Nickname
}
func GetUser(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	var user User
	db.Where("id", uid).First(&user)
	friends := getFriendsById(newUid)
	ctx.JSON(200, gin.H{
		"code":    "200",
		"message": "success",
		"data":    gin.H{"user": user, "friends": friends},
	})
}
func GetUserByUid(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	log.Println(uid, newUid)
	var user User
	db.Where("id", newUid).First(&user)

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}
func UserFriends(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	users := getFriendsById(newUid)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    users,
	})
}
func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")

	var user User
	db.Where("email", email).First(&user)
	friends := getFriendsById(user.Id)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"user":    user,
			"friends": friends,
		},
	})
}
func Register(ctx *gin.Context) {
	var email = ctx.PostForm("email")
	var password = ctx.PostForm("password")
	var verify = ctx.PostForm("verify")
	if verify != "千里江陵一日还" {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "验证错误",
		})
		return
	}
	var user User
	db.Where("email", email).First(&user)
	if user.Id != 0 {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "邮箱已注册",
		})
		return
	}
	data := []byte(password)
	md5Str := fmt.Sprintf("%x", md5.Sum(data))
	pwd := md5Str

	oneUser := RegisterUser{Email: email, Password: pwd, Nickname: email, HeadImg: "https://bulma.io/images/placeholders/96x96.png"}
	result := db.Select("nickname", "password", "email", "head_img").Create(&oneUser)

	//判断插入成功
	if result.RowsAffected != 0 {
		db.Where("email", email).First(&oneUser)
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"uid": oneUser.Id,
			},
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "注册失败",
		})
	}

}

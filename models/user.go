package models

import (
	"app/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var key = "fWsl8sEV4Jw2G!Q9!20Vl*pSyZebqoyr"

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
	Sex      string `gorm:"sex" json:"sex"`
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

//当前用户信息
func GetUserInfo(uid int) User {
	var user User
	db.Model(&User{}).Where("id", uid).First(&user)
	return user
}
func GetOneUser(email string, password string) User {
	newPassword := utils.Md5(password + key)
	var user User
	db.Where("email", email).Where("password", newPassword).First(&user)
	return user
}
func AddUser(email string, password string, sex string) User {
	newPassword := utils.Md5(password + key)
	var res User
	oneUser := RegisterUser{Email: email, Password: newPassword, Nickname: email, HeadImg: "https://chat.daguozhensi.com/images/head_img.png", Sex: sex}
	result := db.Select("nickname", "password", "email", "head_img", "sex").Create(&oneUser)
	if result.RowsAffected != 0 {
		return res
	}
	var one User
	row := db.Model(&User{}).Where("email", email).First(&one)
	if row.Error != nil {
		return res
	}
	return one
}
func UpdateHeadImg(uid int, imgSrc string) bool {
	affected := db.Model(&User{}).Where("id", uid).Update("head_img", imgSrc)
	return affected.RowsAffected != 0
}

//更新昵称
func SetUserNickname(uid int, nickname string) bool {

	affected := db.Model(&User{}).Where("id", uid).Update("nickname", nickname)
	if affected.Error != nil {
		return false
	} else {
		return true
	}

}

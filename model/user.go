package model

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
func GetUser(w http.ResponseWriter, r *http.Request) {
	var uid = r.PostFormValue("uid")
	var user User
	db.Where("id", uid).First(&user)
	var res map[string]interface{}
	res = make(map[string]interface{})

	var data map[string]interface{}
	data = make(map[string]interface{})

	data["user"] = user
	data["friends"] = getFriendsById(1)

	res["code"] = 200
	res["message"] = "success"
	res["data"] = data

	resdata, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(resdata))
}
func GetUserByUid(w http.ResponseWriter, r *http.Request) {
	var uid = r.PostFormValue("uid")
	user_id, err := strconv.Atoi(uid)
	if err != nil {

	}
	var user User
	db.Where("id", user_id).First(&user)
	var res map[string]interface{}
	res = make(map[string]interface{})

	res["code"] = 200
	res["message"] = "success"
	res["data"] = user

	data, err := json.Marshal(res)
	if err != nil {
		panic(err)

	}
	fmt.Fprint(w, string(data))
}
func UserFriends(w http.ResponseWriter, r *http.Request) {
	var uid = r.PostFormValue("uid")
	user_id, err := strconv.Atoi(uid)
	if err != nil {
		return
	}
	var users []User
	users = getFriendsById(user_id)

	var res map[string]interface{}
	res = make(map[string]interface{})

	res["code"] = 200
	res["message"] = "success"
	res["data"] = users
	data, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(data))
}
func Login(w http.ResponseWriter, r *http.Request) {
	var email = r.PostFormValue("email")

	var user User
	db.Where("email", email).First(&user)

	var res map[string]interface{}
	res = make(map[string]interface{})

	var data map[string]interface{}
	data = make(map[string]interface{})

	data["user"] = user
	data["friends"] = getFriendsById(user.Id)

	res["code"] = 200
	res["message"] = "success"
	res["data"] = data

	resdata, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(resdata))
}
func Register(w http.ResponseWriter, r *http.Request) {
	var email = r.PostFormValue("email")
	var password = r.PostFormValue("password")
	var verify = r.PostFormValue("verify")
	if verify != "千里江陵一日还" {
		fmt.Fprint(w, string("{\"code\":0,\"message\":\"验证错误\"}"))
		return
	}
	var user User
	db.Where("email", email).First(&user)
	if user.Id != 0 {
		fmt.Fprint(w, string("{\"code\":0,\"message\":\"邮箱已注册\"}"))
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
		fmt.Println(result)
		fmt.Println(oneUser)
		res := make(map[string]interface{})
		udata := make(map[string]interface{})
		udata["uid"] = oneUser.Id
		res["code"] = 200
		res["message"] = "success"
		res["data"] = udata
		str, err := json.Marshal(res)
		if err != nil {

		}
		fmt.Fprint(w, string(str))
	} else {
		fmt.Fprint(w, string("{\"code\":0,\"message\":\"注册失败\"}"))
	}

}

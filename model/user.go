package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id       int      `gorm:"uid" json:"id"`
	Nickname string   `gorm:"nickname" json:"nickname"`
	Friends  []Friend `gorm:"foreignKey:uid" json:"firends,omitempty"`
}
type OneUser struct {
	Id       int    `gorm:"uid" json:"id"`
	Nickname string `gorm:"nickname" json:"nickname"`
	HeadImg  string `gorm:"head_img" json:"head_img"`
}

func (u *User) TableName() string {
	return "gc_users"
}
func (ou *OneUser) TableName() string {
	return "gc_users"
}
func GetUserName(uid int) string {
	var user User
	db.Where("id=?", uid).First(&user)
	return user.Nickname
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	db.First(&user)
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

package model

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type User struct {
	Id int `gorm:"uid" json:"id"`
	Nickname string `gorm:nickname json:"nickname"`
}

func(u *User) TableName() string {
	return "gc_users"
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	var res map[string]interface{}
	res = make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	res["data"] = users
	data,err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w,string(data))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	db.First(&user)
	var res map[string]interface{}
	res = make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	res["data"] = user
	data,err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w,string(data))
}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("email"))
	var user User
	db.First(&user)
	var res map[string]interface{}
	res = make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	res["data"] = user
	data,err := json.Marshal(res)
	if err != nil {
		panic(err)
		return
	}
	fmt.Fprint(w,string(data))
}
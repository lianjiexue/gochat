package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Mood struct {
	Id      int    `gorm:"id" json:"id"`
	Content string `gorm:"content" json:"content"`
	Uid     int    `gorm:"uid" json:"uid"`
	Time    int    `gorm:"time" json:"time"`
	IsDel   int    `gorm:"is_del" json:"is_del"`
	IsRead  int    `gorm:"is_read" json:"is_read"`
}

func (m *Mood) TableName() string {
	return "gc_mood"
}
func AddMood(w http.ResponseWriter, r *http.Request) {
	var mood Mood
	uid, err := strconv.Atoi(r.FormValue("uid"))
	if err != nil {
		return
	}
	mood.Uid = uid
	mood.Content = r.FormValue("content")
	mood.Time = int(time.Now().Unix())
	db.Save(&mood)
	fmt.Fprintf(w, "{\"code\":200}")
}
func OneMood(w http.ResponseWriter, r *http.Request) {
	var mood Mood
	//读取记录
	db.Where("is_read=0").First(&mood)
	log.Println(mood)
	if mood.Id != 0 {
		//将记录设置为已读
		db.Model(&Mood{}).Where("id", mood.Id).Update("is_read", 1)

		var user OneUser
		db.Where("id", mood.Uid).First(&user)
		res := make(map[string]interface{})
		data := make(map[string]interface{})
		data["user"] = user
		data["mood"] = mood
		res["code"] = 200
		res["data"] = data
		result, err := json.Marshal(res)
		if err != nil {

		}
		fmt.Fprintf(w, string(result))
	} else {
		fmt.Fprintf(w, "{\"code\":0}")
	}

}

func DelMood(w http.ResponseWriter, r *http.Request) {
	db.Model(&Mood{}).Where("id=1").Update("is_del", 1)
	res := make(map[string]interface{})
	res["code"] = 200
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, string(data))
	}
}

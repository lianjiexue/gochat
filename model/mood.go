package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Mood struct {
	Id      int    `gorm:"id"`
	Content string `gorm:"content"`
	Uid     int    `gorm:"uid"`
	Time    int    `gorm:"time"`
	IsDel   int    `gorm:"isdel"`
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
}
func OneMood(w http.ResponseWriter, r *http.Request) {
	var mood Mood
	db.Where("read=0").First(&mood)
	res := make(map[string]interface{})
	res["code"] = 200
	res["data"] = mood
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, string(data))
	}

}

func DelMood(w http.ResponseWriter, r *http.Request) {
	db.Model(&Mood{}).Where("id=1").Update("isdel", 1)
	res := make(map[string]interface{})
	res["code"] = 200
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, string(data))
	}
}

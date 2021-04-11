package model

import (
	"net/http"
	"strconv"
	"time"
)

type Mood struct {
	Id      int    `gorm:"id"`
	Content string `gorm:"content"`
	Uid     int    `gorm:"uid"`
	Time    int    `gorm:"time"`
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

package model

import (
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Id        int    `gorm:"id"`
	MessageId string `gorm:"message_id"`
	FromId    int    `gorm:"from_id"`
	ToId      int    `gorm:"to_id"`
	Cotnent   string `gorm:"content"`
	Time      int    `gorm:"time"`
}

func (m *Message) TableName() string {
	return "gc_message"
}

// 发心情

func AddMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	uid, err := strconv.Atoi(r.FormValue("from_id"))
	if err != nil {
		return
	}
	toid, err := strconv.Atoi(r.FormValue("to_id"))
	if err != nil {

	}
	msg.MessageId = "asdf-asdf-asdf-asdf"
	msg.FromId = uid
	msg.ToId = toid
	msg.Cotnent = r.FormValue("content")
	msg.Time = int(time.Now().Unix())
	db.Save(&msg)
}

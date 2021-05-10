package models

import (
	"log"
	"time"

	"github.com/go-basic/uuid"
)

type MessageMix struct {
	Id        int    `gorm:"id" json:"id"`
	MessageId string `gorm:"message_id" json:"message_id"`
	FromId    int    `gorm:"from_id" json:"from_id"`
	ToId      int    `gorm:"to_id" json:"to_id"`
	Content   string `gorm:"content" json:"content"`
	IsRead    int    `gorm:"is_read" json:"is_read"`
}

type Message struct {
	MessageMix
	Time int `gorm:"time" json:"time"`
}
type MessageNew struct {
	MessageMix
	Datetime string `gorm:"time" json:"datetime"`
}

func (m *Message) TableName() string {
	return "gc_message"
}

// 发新消息
func NewMessage(from_id int, to_id int, content string) MessageNew {
	var msg Message
	var msgnew MessageNew
	msg.MessageId = uuid.New()
	msg.FromId = from_id
	msg.ToId = to_id
	msg.Content = content
	msg.Time = int(time.Now().Unix())
	msg.IsRead = 0
	log.Println(msg)
	affected := db.Model(&Message{}).Create(&msg)
	if affected.Error != nil {
		return msgnew
	}

	msgnew.MessageMix = msg.MessageMix
	msgnew.Datetime = time.Unix(int64(msg.Time), 0).Format("2006.1.2")
	return msgnew
}
func GetFullUnReadMessage(uid int) []Message {
	var message []Message
	db.Model(&Message{}).Where("to_id", uid).Where("is_read", 0).Find(&message)
	db.Model(&Message{}).Where("to_id", uid).Where("is_read", 0).Update("is_read", 1)
	return message
}
func SetReaded(id int) {
	db.Model(&Message{}).Where("id", id).Update("is_read", 1)
}

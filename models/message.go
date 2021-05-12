package models

import (
	"encoding/json"
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

var msgs = make(chan Message, 10)

func writeToRedis() {
	//开一个协程去处理消息
	for {
		select {
		case msg := <-msgs:
			data, _ := json.Marshal(msg)
			//将消息写入redis
			Rdb.LPush(Rdbctx, "messages", data)
			//db.Model(&Message{}).Create(&msg)
		}
	}
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
	msgs <- msg
	go writeToRedis()
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

//从redis中去读取 10个消息批量写入
func SaveMessage() {
	var messages []Message
	log.Println("执行一次批量保存")
	for i := 0; i < 10; i++ {
		rdbcmd := Rdb.LPop(Rdbctx, "messages")
		result, err := rdbcmd.Result()
		if err != nil {
			var oneMessage Message
			json.Unmarshal([]byte(result), &oneMessage)
			messages = append(messages, oneMessage)
		}

	}
	//批量写入数据库
	if len(messages) > 0 {
		db.Create(messages)
	}

}

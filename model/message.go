package model

import (
	"app/socket"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

type MessageMix struct {
	Id        int    `gorm:"id" json:"id"`
	MessageId string `gorm:"message_id" json:"message_id"`
	FromId    int    `gorm:"from_id" json:"from_id"`
	ToId      int    `gorm:"to_id" json:"to_id"`
	Content   string `gorm:"content" json:"content"`
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
func NewMessage(serve *socket.Serve, ctx *gin.Context) {
	var msg Message
	uid, err := strconv.Atoi(ctx.PostForm("from_id"))
	if err != nil {
		return
	}
	toid, err := strconv.Atoi(ctx.PostForm("to_id"))
	if err != nil {
		return
	}
	msg.MessageId = uuid.New()
	msg.FromId = uid
	msg.ToId = toid
	msg.Content = ctx.PostForm("content")
	msg.Time = int(time.Now().Unix())
	db.Save(&msg)
	var msgnew MessageNew
	msgnew.MessageMix = msg.MessageMix
	msgnew.Datetime = time.Unix(int64(msg.Time), 0).Format("2006.1.2")

	log.Println(msgnew)
	is_socket := SendMsg(msgnew, serve)

	if is_socket {
		ctx.JSON(200, gin.H{
			"code": 200,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}
func SendMsg(msg MessageNew, serve *socket.Serve) bool {
	cli := serve.GetClinet(msg.ToId)
	data := make(map[string]interface{})
	data["type"] = "message"
	data["msg"] = msg
	result, err := json.Marshal(data)
	if err != nil {
		cli.Conn.WriteMessage(1, []byte("发送失败"))
		return false
	}

	if cli.Conn != nil {
		cli.Conn.WriteMessage(1, []byte(result))
	} else {
		return false
	}

	return true
}

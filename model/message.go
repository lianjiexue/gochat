package model

import (
	"app/socket"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
func NewMessage(serve *socket.Serve, w http.ResponseWriter, r *http.Request) {
	var msg Message
	uid, err := strconv.Atoi(r.FormValue("from_id"))
	if err != nil {
		return
	}
	toid, err := strconv.Atoi(r.FormValue("to_id"))
	if err != nil {

	}
	msg.MessageId = uuid.New()
	msg.FromId = uid
	msg.ToId = toid
	msg.Content = r.FormValue("content")
	msg.Time = int(time.Now().Unix())
	db.Save(&msg)
	var msgnew MessageNew
	msgnew.MessageId = msg.MessageId
	msgnew.Content = msg.Content
	msgnew.FromId = msg.FromId
	msgnew.ToId = msg.ToId
	msgnew.Content = msg.Content
	msgnew.Datetime = time.Unix(int64(msg.Time), 0).Format("2006.1.2")

	log.Println(msgnew)
	is_socket := SendMsg(msgnew, serve)

	if is_socket == true {
		fmt.Fprintf(w, "{\"code\":200}")
	} else {
		fmt.Fprintf(w, "{\"code\":0}")
	}
}
func SendMsg(msg MessageNew, serve *socket.Serve) bool {
	data := make(map[string]interface{})
	data["type"] = "message"
	data["msg"] = msg
	result, err := json.Marshal(data)
	if err != nil {

	}
	log.Println(msg.ToId, msg.Content)
	cli := serve.GetClinet(msg.ToId)
	log.Println(cli)
	if cli.Conn != nil {
		log.Println("对象客户端")
		log.Println(cli)
		cli.Conn.WriteMessage(1, []byte(result))
	} else {
		return false
	}

	return true
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conn *websocket.Conn

// 1.bind  2.online  3.message
type Message struct {
	MessageType string `json:"type"`
	MessageId   string `json:"messageid"`
	FromId      string `json:"fromid"`
	ToId        string `json:"toid"`
	Content     string `json:"content"`
}

//读消息
func readLoop(conn *websocket.Conn) {
	messagetype, data, err := conn.ReadMessage()

	if err != nil {
		log.Println("断开链接")
		for client, _ := range clients {
			if client.Conn == conn {
				delete(clients, client)
			}
		}
		return
	}
	//写数据
	var message Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		conn.WriteMessage(messagetype, []byte("发送失败"))
		return
	}
	if message.MessageType != "online" || message.MessageType != "bind" || message.MessageType != " message" {
		conn.WriteMessage(messagetype, []byte("格式错误"))
		return
	}
	switch message.MessageType {
	case "pong":
		return
	case "bind":
		client := new(Client)
		client.Uid = message.FromId
		client.Conn = conn
		clients[Conn] = client

	case "online":
		conn.WriteMessage(messagetype, onlineUser())
	case "message":
		var Uid string
		var client = new(Client)
		client = getOneClient()
		Uid = client.ToId
		client = getOneClient(Uid)
		client.Conn.WriteMessage(messagetype, data)
	default:
		return
	}
	err = conn.WriteMessage(messagetype, data)

	if err != nil {
		log.Println("发送失败")
		return
	}
}
func service(w http.ResponseWriter, r *http.Request) {
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("链接失败")
		return
	}
	log.Println(time.Now())

	for {
		readLoop(conn)
	}
}

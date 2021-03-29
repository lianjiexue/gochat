package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// 1.bind  2.online  3.message
type Message struct {
	MessageType string `json:"type"`
	MessageId   string `json:"messageid"`
	FromId      string `json:"fromid"`
	ToId        string `json:"toid"`
	Content     string `json:"content"`
}

func service(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("链接失败")
		return
	}
	log.Println(time.Now())

	for {
		//读数据
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
		case "bind":
			client := new(Client)
			client.Uid = message.FromId
			client.Conn = conn
			clients[client] = true
			return

		case "online":
			conn.WriteMessage(messagetype, onlineUser())
			return

		case "message":
			var nowUid string

			for client, _ := range clients {

				if client.Conn == conn {
					nowUid = client.Uid
				}
				if client.Uid == message.ToId {
					var msg = new(Message)
					msg.MessageType = "message"
					msg.FromId = nowUid
					msg.Content = message.Content
					data, err := json.Marshal(msg)
					if err != nil {
						return
					}
					client.Conn.WriteMessage(messagetype, data)
				}
			}

		}
		err = conn.WriteMessage(messagetype, data)

		if err != nil {
			log.Println("发送失败")
			return
		}
	}
}

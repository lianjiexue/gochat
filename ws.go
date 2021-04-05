package main

import (
	"app/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conn *websocket.Conn

// 1.bind  2.online  3.message
type Message struct {
	Type      string `json:"type"`
	MessageId string `json:"messageid"`
	FromId    int    `json:"fromid"`
	ToId      int    `json:"toid"`
	Content   string `json:"content"`
}
type OnlineUser struct {
	Uid      int    `json:"uid"`
	NickName string `json:"nickname"`
}

func online() []byte {
	var res map[string]interface{}
	res = make(map[string]interface{})
	var onlineUsers []OnlineUser
	for _, client := range clients {
		var u OnlineUser
		u.Uid = client.Uid
		u.NickName = client.NickName
		onlineUsers = append(onlineUsers, u)
	}
	res["code"] = 200
	res["message"] = "success"
	res["users"] = onlineUsers
	users, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return users
}

//读消息
func readLoop(conn *websocket.Conn) {
	messagetype, data, err := conn.ReadMessage()

	//写数据
	var message Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		conn.WriteMessage(messagetype, []byte("发送失败"))
		log.Println(err)
		return
	} else {
		log.Println(message.Type)
	}
	if message.Type != "online" && message.Type != "bind" && message.Type != "message" && message.Type != "pong" {
		conn.WriteMessage(messagetype, []byte("格式错误"))
		return
	}
	switch message.Type {
	case "pong":
		return
	case "bind":
		var client Client
		if message.FromId == 0 {
			conn.WriteMessage(messagetype, []byte("格式错误"))
			return
		}
		client.Uid = message.FromId
		client.NickName = model.GetUserName(client.Uid)
		client.Conn = conn
		clients[conn] = client

	case "online":
		conn.WriteMessage(messagetype, online())
	case "message":
		var Uid int
		Uid = message.ToId
		var client Client
		client = getOneClient(Uid)
		log.Println("对象客户端")
		log.Println(client)
		if client.Conn != nil {
			client.Conn.WriteMessage(messagetype, data)
		} else {
			conn.WriteMessage(messagetype, []byte("对象已离线"))
		}
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
	defer conn.Close()
	if err != nil {
		log.Println("链接失败")
		return
	}
	log.Println(clients)

	for {

		readLoop(conn)
	}
}

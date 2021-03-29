package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/websocket"
	"time"
	"app/db"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
type Client struct {
	Conn  *websocket.Conn
	Uid string
}
// 1.bind  2.online  3.message
type Message struct {
	MessageType string  `json:"type"`
	MessageId string  `json:"messageid"`
	FromId   string   `json:"fromid"`
	ToId     string   `json:"toid"`
	Content  string   `json:"content"`	
}
var err error
var clients map[*Client]bool
func init(){
	log.Println("runing 127.0.0.1:8080")
	clients = make(map[*Client]bool)
}
func onlineUser()[]byte {
	var res map[string]interface{}
		res = make(map[string]interface{})
		res["code"] = 200
		res["message"] = "success"
		res["users"] = clients
	users,err := json.Marshal(res)
	if err != nil{
		panic(err)
	}
	return users
}
func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"index.html")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
		var res map[string]interface{}
		res = make(map[string]interface{})
		res["code"] = 200
		res["message"] = "success"
		data,err := json.Marshal(res)

		if err != nil {}
		fmt.Fprint(w,data)
	})
	http.HandleFunc("/api/users", db.GetUsers)
	http.HandleFunc("/api/user", db.GetUser)
	http.HandleFunc("/api/login", db.Login)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		conn,err := upgrader.Upgrade(w,r,nil)
		if err != nil{
			log.Println("链接失败")
			return
		}
		log.Println(time.Now())
		
		for {
			//读数据
			messagetype,data,err := conn.ReadMessage()
			
			if err != nil{
				log.Println("断开链接")
				for client,_ := range clients{
					if client.Conn == conn {
						delete(clients,client)
					}
				}
				return
			}
			//写数据
			var message Message
			err = json.Unmarshal(data,&message)
			if err != nil{
				conn.WriteMessage(messagetype,[]byte("发送失败"))
				return
			}
			if message.MessageType != "online" || message.MessageType != "bind" || message.MessageType != " message"{
				conn.WriteMessage(messagetype,[]byte("格式错误"))
				return
			}
			switch (message.MessageType) {
				case "bind" :
					client := new(Client)
					client.Uid = message.FromId
					client.Conn = conn
					clients[client] = true
					return
				;

				case "online":
					conn.WriteMessage(messagetype,onlineUser())
					return
				;
			case "message":
				var nowUid string

				for client,_ := range clients{
					
					if client.Conn == conn {
						nowUid = client.Uid
					}
					if client.Uid == message.ToId {
						var msg = new(Message)
							msg.MessageType = "message"
							msg.FromId = nowUid
							msg.Content = message.Content
						data,err := json.Marshal(msg)
						if err != nil{
								return
						}
						client.Conn.WriteMessage(messagetype,data)
					}
				}
				;
			}
			err = conn.WriteMessage(messagetype,data)

			if err != nil{
				log.Println("发送失败")
				return
			}
		}

	})
	http.ListenAndServe(":8080",nil)
}

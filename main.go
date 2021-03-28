package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/websocket"
	"time"
	"math/rand"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
type Client struct {
	Conn  *websocket.Conn
	Uid int
	Username string
}
type Message struct {
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
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		conn,err := upgrader.Upgrade(w,r,nil)
		if err != nil{
			return
		}
		log.Println(time.Now())
		client := new(Client)
		client.Uid = rand.Intn(1000)
		client.Conn = conn
		clients[client] = true
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
				delete(clients,client)
				return
			}

			log.Println(len(clients))
			log.Println(client)
			//写数据
			var message Message
			err = json.Unmarshal(data,&message)
			if err != nil{
				conn.WriteMessage(messagetype,[]byte("发送失败"))
				return
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

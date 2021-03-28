package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
var err error
func init(){
	log.Println("runing 127.0.0.1:8080")
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
		if err != nil{}

		for {
			//读数据
			messagetype,data,err := conn.ReadMessage()
			
			if err != nil{}

			//写数据

			err = conn.WriteMessage(messagetype,data)

			if err != nil{}
		}

	})
	http.ListenAndServe(":8080",nil)
}

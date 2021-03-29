package main

import (
	"app/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var err error

func init() {
	log.Println("runing 127.0.0.1:8080")

}
func onlineUser() []byte {
	var res map[string]interface{}
	res = make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	res["users"] = clients
	users, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return users
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		var res map[string]interface{}
		res = make(map[string]interface{})
		res["code"] = 200
		res["message"] = "success"
		data, err := json.Marshal(res)

		if err != nil {
		}
		fmt.Fprint(w, data)
	})
	http.HandleFunc("/api/users", db.GetUsers)
	http.HandleFunc("/api/user", db.GetUser)
	http.HandleFunc("/api/login", db.Login)
	//ws服务
	http.HandleFunc("/ws", service)
	http.ListenAndServe(":8080", nil)
}

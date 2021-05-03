package main

import (
	"app/model"
	"app/socket"
	"log"
	"net/http"
)

var err error

func init() {
	log.Println("runing 127.0.0.1:8008")

}

func main() {

	serve := &socket.Serve{Clients: make(map[string]*socket.Client), Messages: make(chan []byte), On: make(chan *socket.Client), Off: make(chan *socket.Client)}
	go serve.Run()
	http.HandleFunc("/api/user", model.GetUser)
	http.HandleFunc("/api/login", model.Login)
	http.HandleFunc("/api/register", model.Register)

	http.HandleFunc("/api/user/friends", model.UserFriends)
	http.HandleFunc("/api/message/new", func(w http.ResponseWriter, r *http.Request) {
		model.NewMessage(serve, w, r)
	})
	//个人
	http.HandleFunc("/api/user/info", model.GetUserByUid)
	//心情
	http.HandleFunc("/api/mood/add", model.AddMood)
	http.HandleFunc("/api/mood/one", model.OneMood)
	http.HandleFunc("/api/mood/del", model.DelMood)

	//ws服务
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.Ws(serve, w, r)
	})
	http.ListenAndServe(":8008", nil)
}

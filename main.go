package main

import (
	"app/model"
	"log"
	"net/http"
)

var err error

func init() {
	log.Println("runing 127.0.0.1:8081")

}

func main() {

	serve := &Serve{Clients: make(map[string]*Client), Messages: make(chan []byte), On: make(chan *Client), Off: make(chan *Client)}
	go serve.run()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/api/user", model.GetUser)
	http.HandleFunc("/api/login", model.Login)
	http.HandleFunc("/api/user/friends", model.UserFriends)
	//ws服务
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws(serve, w, r)
	})
	http.ListenAndServe(":8081", nil)
}

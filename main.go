package main

import (
	"app/model"
	"log"
	"net/http"
)

var err error

func init() {
	log.Println("runing 127.0.0.1:8080")

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/api/user", model.GetUser)
	http.HandleFunc("/api/login", model.Login)
	http.HandleFunc("/api/user/friends", model.UserFriends)
	//ws服务
	http.HandleFunc("/ws", service)
	http.ListenAndServe(":8080", nil)
}

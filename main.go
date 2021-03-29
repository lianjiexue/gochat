package main

import (
	"app/model"
	"encoding/json"
	"log"
	"net/http"
)

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
	http.HandleFunc("/api/users", model.GetUsers)
	http.HandleFunc("/api/user", model.GetUser)
	http.HandleFunc("/api/login", model.Login)
	//ws服务
	http.HandleFunc("/ws", service)
	http.ListenAndServe(":8080", nil)
}

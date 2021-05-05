package main

import (
	"app/model"
	"app/socket"
	"log"

	"github.com/gin-gonic/gin"
)

var err error

func init() {
	log.Println("runing 127.0.0.1:8008")

}

func main() {

	serve := &socket.Serve{Clients: make(map[string]*socket.Client), Messages: make(chan []byte), On: make(chan *socket.Client), Off: make(chan *socket.Client)}
	go serve.Run()

	router := gin.Default()

	router.POST("/api/user", model.GetUser)
	router.POST("/api/login", model.Login)
	router.POST("/api/register", model.Register)
	//个人
	router.POST("/api/user/info", model.GetUserByUid)
	router.POST("/api/user/follow", model.GetUserFollow)
	router.POST("/api/user/friends", model.UserFriends)
	router.POST("/api/message/new", func(ctx *gin.Context) {
		model.NewMessage(serve, ctx)
	})
	//心情
	router.POST("/api/mood/add", model.AddMood)
	router.POST("/api/mood/one", model.OneMood)
	router.POST("/api/mood/del", model.DelMood)

	//好友
	router.POST("/api/friend/follow", model.Follow)
	router.POST("/api/friend/unfollow", model.UnFollow)
	//ws服务
	router.GET("/ws", func(ctx *gin.Context) {
		socket.Ws(serve, ctx)
	})
	router.Run(":8008")
}

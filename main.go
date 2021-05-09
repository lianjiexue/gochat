package main

import (
	"app/controller"
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
	router.MaxMultipartMemory = 2 << 20

	router.POST("/api/login", controller.Login)
	router.POST("/api/register", controller.Register)
	//个人
	router.POST("/api/user/info", controller.UserInfo)
	router.POST("/api/user/follow", controller.UserFollow) //返回当前人物信息，并且返回是否已经关注
	router.POST("/api/user/friends", controller.MyFriends)
	router.POST("/api/message/new", func(ctx *gin.Context) {
		controller.NewMessage(serve, ctx)
	})
	router.POST("/api/user/update", controller.UpdateNickname)
	router.POST("/api/user/upload", controller.UpdateHeadImg)
	//心情
	router.POST("/api/mood/add", controller.AddMood)
	router.POST("/api/mood/one", controller.OneMood)
	router.POST("/api/mood/del", controller.DelMood)

	//好友
	router.POST("/api/friend/follow", controller.Follow)
	router.POST("/api/friend/unfollow", controller.Unfollow)
	//ws服务
	router.GET("/ws", func(ctx *gin.Context) {
		socket.Ws(serve, ctx)
	})
	router.Run(":8008")
}

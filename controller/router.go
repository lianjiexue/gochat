package controller

import (
	"app/socket"

	"github.com/gin-gonic/gin"
)

func Run() {
	serve := &socket.Serve{Clients: make(map[string]*socket.Client), Messages: make(chan []byte), On: make(chan *socket.Client), Off: make(chan *socket.Client)}
	go serve.Run()

	router := gin.Default()
	router.MaxMultipartMemory = 2 << 20

	router.POST("/api/login", Login)
	router.POST("/api/register", Register)
	//个人
	router.POST("/api/user/info", UserInfo)
	router.POST("/api/user/follow", UserFollow) //返回当前人物信息，并且返回是否已经关注
	router.POST("/api/user/friends", MyFriends)
	router.POST("/api/message/new", func(ctx *gin.Context) {
		NewMessage(serve, ctx)
	})
	router.POST("/api/user/update", UpdateNickname)
	router.POST("/api/user/upload", UpdateHeadImg)
	//心情
	router.POST("/api/mood/add", AddMood)
	router.POST("/api/mood/one", OneMood)
	router.POST("/api/mood/del", DelMood)
	router.POST("/api/mood/list", MoodsLists)

	//好友
	router.POST("/api/friend/follow", Follow)
	router.POST("/api/friend/unfollow", Unfollow)
	//ws服务
	router.GET("/ws", func(ctx *gin.Context) {
		socket.Ws(serve, ctx)
	})
	router.Run(":8008")
}

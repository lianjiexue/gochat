package router

import (
	"app/controllers"
	"app/socket"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	serve := &socket.Serve{Clients: make(map[string]*socket.Client), Messages: make(chan []byte), On: make(chan *socket.Client), Off: make(chan *socket.Client)}
	go serve.Run()
	//开一个定时任务,实现消息的批量写入数据库
	go func() {
		Timer := time.NewTicker(60 * time.Second)
		select {
		case <-Timer.C:
			controllers.SaveMessage()
		}
	}()
	router := gin.Default()
	router.MaxMultipartMemory = 2 << 20

	router.POST("/api/login", controllers.Login)
	router.POST("/api/register", controllers.Register)
	//个人
	router.POST("/api/user/info", controllers.UserInfo)
	router.POST("/api/user/follow", controllers.UserFollow) //返回当前人物信息，并且返回是否已经关注
	router.POST("/api/user/friends", controllers.MyFriends)
	router.POST("/api/message/new", func(ctx *gin.Context) {
		controllers.NewMessage(serve, ctx)
	})
	router.POST("/api/message/unread", controllers.FullUnRead)
	router.POST("/api/user/update", controllers.UpdateNickname)
	router.POST("/api/user/upload", controllers.UpdateHeadImg)
	//心情
	router.POST("/api/mood/add", controllers.AddMood)
	router.POST("/api/mood/one", controllers.OneMood)
	router.POST("/api/mood/del", controllers.DelMood)
	router.POST("/api/mood/list", controllers.MoodsLists)

	//好友
	router.POST("/api/friend/follow", controllers.Follow)
	router.POST("/api/friend/unfollow", controllers.Unfollow)
	//ws服务
	router.GET("/ws", func(ctx *gin.Context) {
		socket.Ws(serve, ctx)
	})
	router.Run(":8008")
}

package controller

import (
	"app/model"
	"app/socket"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewMessage(serve *socket.Serve, ctx *gin.Context) {
	from_id := ctx.PostForm("from_id")
	newFromId, _ := strconv.Atoi(from_id)
	to_id := ctx.PostForm("to_id")
	ToId, _ := strconv.Atoi(to_id)

	content := ctx.PostForm("content")

	msgnew := model.NewMessage(newFromId, ToId, content)
	log.Println(msgnew)
	if msgnew.Id != 0 {
		is_send := SendMsg(msgnew, serve)
		if is_send {
			ctx.JSON(200, gin.H{
				"code":    200,
				"message": "success",
			})
		} else {
			ctx.JSON(200, gin.H{
				"code":    0,
				"message": "发送失败",
			})
		}

	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "消息保存失败",
		})
	}

}
func SendMsg(msg model.MessageNew, serve *socket.Serve) bool {
	cli := serve.GetClinet(msg.ToId)
	data := make(map[string]interface{})
	data["type"] = "message"
	data["msg"] = msg
	result, err := json.Marshal(data)
	if err != nil {
		cli.Conn.WriteMessage(1, []byte("发送失败"))
		return false
	}

	if cli.Conn != nil {
		cli.Conn.WriteMessage(1, []byte(result))
	} else {
		return false
	}

	return true
}

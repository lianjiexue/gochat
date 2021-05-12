package controllers

import (
	"app/models"
	"app/socket"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewMessage(serve *socket.Serve, ctx *gin.Context) {
	from_id := ctx.PostForm("from_id")
	newFromId, _ := strconv.Atoi(from_id)
	to_id := ctx.PostForm("to_id")
	ToId, _ := strconv.Atoi(to_id)

	content := ctx.PostForm("content")

	msgnew := models.NewMessage(newFromId, ToId, content)

	if msgnew.ToId != 0 {
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

func SendMsg(msg models.MessageNew, serve *socket.Serve) bool {
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
		// setMessageRead(msg.Id)
	} else {
		return false
	}

	return true
}

//返回未读消息
func FullUnRead(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	newUid, _ := strconv.Atoi(uid)
	messages := models.GetFullUnReadMessage(newUid)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    messages,
	})
}
func SaveMessage() {
	models.SaveMessage()
}

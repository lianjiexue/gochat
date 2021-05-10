package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Serve struct {
	Clients  map[string]*Client
	Messages chan []byte
	On       chan *Client
	Off      chan *Client
}

type Client struct {
	Conn     *websocket.Conn
	Id       string
	Uid      int
	Serve    *Serve
	NickName string
	Send     chan []byte
}

type Message struct {
	Type      string `json:"type"`
	MessageId string `json:"messageid"`
	FromId    int    `json:"fromid"`
	ToId      int    `json:"toid"`
	Content   string `json:"content"`
}

func Ws(serve *Serve, ctx *gin.Context) {
	conn, error := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if error != nil {
		log.Println("未链接")
		log.Print(error)
		return
	}
	guid := xid.New()
	id := guid.String()
	//生成一个用户
	client := &Client{Conn: conn, Serve: serve, Id: id, Send: make(chan []byte)}
	//注册到websocket服务中
	serve.On <- client
	go client.Done()
}
func (c *Client) Done() {
	defer func() {
		c.Serve.Off <- c
		c.Conn.Close()
	}()
	for {
		dataType, data, err := c.Conn.ReadMessage()
		//判断消息是否能获取到
		if err != nil {
			log.Println("error")
			return
		}
		var message Message
		err = json.Unmarshal(data, &message)
		if err != nil {
			log.Println("消息解析失败")
			return
		}
		if message.Type != "online" && message.Type != "bind" && message.Type != "pong" {
			c.Conn.WriteMessage(dataType, []byte("{\"type\":\"online\"}"))
			return
		}
		//分类消息处理
		switch message.Type {
		case "ping":
			c.Conn.WriteMessage(dataType, []byte("{\"type\":\"ping\"}"))
		case "bind":
			c.Uid = message.FromId
			c.Conn.WriteMessage(dataType, []byte("{\"type\":\"bind\"}"))
		}
	}
}
func (s *Serve) GetClinet(uid int) *Client {
	cli := new(Client)
	for _, client := range s.Clients {
		if client.Uid == uid {
			cli = client
			break
		}
	}
	return cli
}
func (s *Serve) Run() {
	for {
		select {
		case client := <-s.On:
			// 用户接入
			s.Clients[client.Id] = client
		case client := <-s.Off:
			delete(s.Clients, client.Id)
		}
	}
}

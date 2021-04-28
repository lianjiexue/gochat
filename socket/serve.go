package socket

import (
	"log"
	"net/http"

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

func Ws(serve *Serve, w http.ResponseWriter, r *http.Request) {
	conn, error := upgrader.Upgrade(w, r, nil)
	if error != nil {
		log.Println("未链接")
	}
	guid := xid.New()
	id := guid.String()
	//生成一个用户
	client := &Client{Conn: conn, Serve: serve, Id: id, Send: make(chan []byte)}
	//写入到服务中
	serve.On <- client
	go client.ReadMsg()
	go client.WriteMsg()
}
func (s *Serve) GetClinet(uid int) *Client {
	cli := new(Client)
	log.Println(s.Clients)
	for _, client := range s.Clients {
		//获取到客户端
		log.Println("获取到客户端", client)
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
			log.Println(client, "用户接入")
		case client := <-s.Off:
			log.Println("用户离开", client.Id)
			//用户离开
			delete(s.Clients, client.Id)
			close(client.Send)
			log.Println(client, "用户离开")
		case message := <-s.Messages:
			// 判断type 进行处理
			// done(message)
			for _, client := range s.Clients {
				select {
				case client.Send <- message:
				}
			}
		}
	}
}

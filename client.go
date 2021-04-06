package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Id       string
	Uid      int
	Serve    *Serve
	NickName string
	Send     chan []byte
}

// 1.bind  2.online  3.message,4.ground
type Message struct {
	Type      string `json:"type"`
	MessageId string `json:"messageid"`
	FromId    int    `json:"fromid"`
	ToId      int    `json:"toid"`
	Content   string `json:"content"`
}

func done(client *Client, data []byte) {
	//写数据
	var message Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		client.Conn.WriteMessage(1, []byte("发送失败"))
		log.Println(err)
		return
	} else {
		log.Println(message.Type)
	}
	if message.Type != "online" && message.Type != "bind" && message.Type != "message" && message.Type != "pong" {
		client.Conn.WriteMessage(1, []byte("格式错误"))
		return
	}
	//分类消息处理
	switch message.Type {
	case "pong":
		return
	case "bind":
		client.Uid = message.FromId
	case "message":
		var Uid int
		Uid = message.ToId
		log.Println(Uid)
		cli := client.Serve.getClinet(Uid)
		log.Println(cli)
		if cli.Conn != nil {
			log.Println("对象客户端")
			log.Println(cli)
			cli.Conn.WriteMessage(1, []byte(message.Content))
		} else {
			client.Conn.WriteMessage(1, []byte("用户已离线"))
		}

	case "broadcast":
		//群发
		client.Serve.Messages <- data
	}
}

func (c *Client) ReadMsg() {
	defer func() {
		c.Serve.Off <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("获取消息失败")
			return
		}
		done(c, message)
	}
}

func (c *Client) WriteMsg() {
	defer func() {
		c.Serve.Off <- c
		c.Conn.Close()
	}()

	for {
		select {
		case message := <-c.Send:
			c.Conn.WriteMessage(1, message)
		}

	}
}

package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Uid      int
	NickName string
}

var clients map[*websocket.Conn]Client

func init() {
	clients = make(map[*websocket.Conn]Client)
}

func getOneClient(uid int) Client {
	var c Client
	for _, client := range clients {
		if client.Uid == uid {
			return client
		}
		break
	}
	return c
}

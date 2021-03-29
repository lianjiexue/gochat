package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Uid  string
	Username string
}

var clients map[*websocket.Conn]Client
func init() {
	clients = make(map[*Client]bool)
}

func getOneClient(uid int) Client{
	for client,_ := range clients {
		if client.Uid == uid {
			return client
		}
		break
	}
}
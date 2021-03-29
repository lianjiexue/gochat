package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Uid  string
}

var clients map[*Client]bool

func init() {
	clients = make(map[*Client]bool)
}

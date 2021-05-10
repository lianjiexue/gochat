package main

import (
	"app/router"
	"log"
)

func init() {
	log.Println("runing 127.0.0.1:8008")
}

func main() {
	router.Run()
}

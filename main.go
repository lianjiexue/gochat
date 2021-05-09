package main

import (
	"app/controller"
	"log"
)

func init() {
	log.Println("runing 127.0.0.1:8008")
}

func main() {
	controller.Run()
}

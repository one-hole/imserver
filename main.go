package main

import (
	"log"
	"os"

	"gitee.com/odd-socket/sockets"
	"gitee.com/odd-socket/sources"
)

func init() {
	log.SetOutput(os.Stdout)
}

var (
	forever = make(chan bool)
)

func main() {
	manager := sockets.NewManger()
	go sources.Run(manager)
	go sockets.Run(manager)
	go manager.Exec()
	<-forever
}

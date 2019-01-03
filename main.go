package main

import (
	"gitee.com/odd-socket/sockets"
	"gitee.com/odd-socket/sources"
)

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

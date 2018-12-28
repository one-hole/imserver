package main

import (
	"gitee.com/odd-socket/sockets"
)

var (
	forever = make(chan bool)
)

func main() {
	manager := sockets.NewManger()
	defer manager.Close()
	go manager.Exec()
	go sockets.Run(manager)
	<-forever
}

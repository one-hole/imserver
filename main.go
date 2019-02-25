package main

import (
	"log"
	"os"

	"github.com/w-zengtao/socket-server/api"
	"github.com/w-zengtao/socket-server/sockets"
	"github.com/w-zengtao/socket-server/sources"
)

func init() {
	log.SetOutput(os.Stdout)
}

var (
	forever = make(chan bool)
)

func main() {
	manager := sockets.NewManger()
	go api.Run(manager)
	go sources.Run(manager)
	go manager.Exec()
	<-forever
}

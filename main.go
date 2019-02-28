package main

import (
	"log"
	"os"

	"github.com/w-zengtao/socket-server/variable"

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
	variable.Managers = append(variable.Managers, manager)
	go api.Run(manager)
	go sources.RunRabbit(manager)
	go sources.RunRedis(manager)
	go manager.Exec()

	<-forever
}

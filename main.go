package main

import (
	"log"
	"os"

	"github.com/w-zengtao/socket-server/api"
	"github.com/w-zengtao/socket-server/models"
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
	manager := newManager("")

	go api.Run(manager)
	go sources.RunRabbit(manager)
	go sources.RunRedis(manager)
	go manager.Exec()

	<-forever
}

func newManager(name string) *sockets.ClientManager {
	if name == "" {
		name = "default"
	}
	manager := sockets.NewManger()
	models.Managers[name] = manager
	return manager
}

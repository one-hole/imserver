package main

import (
	"log"
	"os"

	"github.com/one-hole/imserver/api"
	"github.com/one-hole/imserver/models"
	"github.com/one-hole/imserver/sockets"
	"github.com/one-hole/imserver/sources"
)

func init() {
	log.SetOutput(os.Stdout)
}

var (
	forever = make(chan bool)
)

func main() {

	go api.Run()

	execManager(newManager("default"), "", "")
	execManager(newManager("csgo"), "aiesports-csgo-websocket", "")
	execManager(newManager("dota2"), "aiesports-dota2-websocket", "")

	models.Init()

	<-forever
}

func newManager(name string) *sockets.ClientManager {
	if name == "" {
		name = "default"
	}
	manager := sockets.NewManger()
	sockets.Managers[name] = manager
	return manager
}

func execManager(m *sockets.ClientManager, redisChannel string, rabbitRouteKey string) error {
	if "" != redisChannel {
		go sources.RunRedis(m, redisChannel)
	}

	if "" != rabbitRouteKey {
		go sources.RunRabbit(m, rabbitRouteKey)
	}

	go m.Exec()

	return nil
}

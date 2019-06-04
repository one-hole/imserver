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

	runManager(newManager("default"), "rw-hz-odds-routing")
	//runManager(newManager("tenants"), "rw-hk-tenants-routing")

	models.Init()

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

func runManager(m *sockets.ClientManager, rabbitRouteKey string) {
	go sources.RunRedis(m)
	// go sources.RunRabbit(m, rabbitRouteKey)
	go m.Exec()
}

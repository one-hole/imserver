package models

import (
	"github.com/one-hole/imserver/sockets"
)

// Managers - 用 name 来区分 manager, 不是很好复用连接 后期还是需要引入 Room 的概念
var (
	Managers = make(map[string]*sockets.ClientManager)
)

// DefaultManager - as it looks
func DefaultManager() *sockets.ClientManager {
	return Managers["default"]
}

func ManagerByName(name string) *sockets.ClientManager {
	if i, ok := Managers[name]; ok {
		return i
	}
	return DefaultManager()
}

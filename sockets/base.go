package sockets

import "github.com/one-hole/imserver/models"

// Managers - 用 name 来区分 manager, 不是很好复用连接 后期还是需要引入 Room 的概念
var (
	Managers = make(map[string]*ClientManager)
)

// DefaultManager - as it looks
func DefaultManager() *ClientManager {
	return Managers["default"]
}

// ManagerByName find manager by name
func ManagerByName(name string) *ClientManager {
	if i, ok := Managers[name]; ok {
		return i
	}
	return DefaultManager()
}

// verifyConn 判断长连接是否该断开 (false 表示需要断开 & true 表示继续维持)
func verifyConn(host string) bool {
	server := &models.Server{}
	if err := server.RecordByHost(host); err != nil {
		return false
	}

	return server.Valid()
}

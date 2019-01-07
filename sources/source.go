package sources

/*
	I want to Interface the sources
	Each source just need to WriteToManager
*/

import (
	"gitee.com/odd-socket/sockets"
)

// Source defines the interface of datasource
type Source interface {
	WriteToManager(m *sockets.ClientManager)
}

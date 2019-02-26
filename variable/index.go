package variable

import "github.com/w-zengtao/socket-server/sockets"

var (
	Managers = make([]*sockets.ClientManager, 0, 10)
)

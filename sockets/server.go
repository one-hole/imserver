package sockets

import (
	"fmt"
	"net/http"

	"gitee.com/odd-socket/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 这里是要帮助 "客户端" 创建连接 & 实例化 Client & 注册该 Client
func serveWs(manager *ClientManager, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	utils.FailOnError(err, "Failed to upgrade a conn")

	client := &Client{
		conn:     conn,
		messages: make(chan []byte, 256),
	}

	manager.register <- client

	go client.readMessageFromClient()
	go client.writeMessageToClient()
}

// Run function runs the loop for the websocket server
func Run(manager *ClientManager) {
	fmt.Println("this is the sockets server")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(manager, w, r)
	})
	http.ListenAndServe(":8080", nil)
}

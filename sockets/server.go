package sockets

import (
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/websocket"
	"github.com/w-zengtao/socket-server/config"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 这里是要帮助 "客户端" 创建连接 & 实例化 Client & 注册该 Client
func serveWs(manager *ClientManager, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		manager:  manager,
		conn:     conn,
		messages: make(chan []byte, 256),
	}

	manager.register <- client

	go client.readMessageFromClient()
	go client.writeMessageToClient()
}

// Run function runs the loop for the websocket server
func Run(manager *ClientManager) {
	http.HandleFunc(fmt.Sprintf("/%s", config.Instance().Socket.Path), func(w http.ResponseWriter, r *http.Request) {
		serveWs(manager, w, r)
	})
	http.ListenAndServe(fmt.Sprintf(":%s", config.Instance().Socket.Port), nil)
}

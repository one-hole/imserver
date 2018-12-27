package sockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(manager *ClientManager, w http.ResponseWriter, r *http.Request) {

}

// Run function runs the loop for the websocket server
func Run(manager *ClientManager) {
	fmt.Println("this is the sockets server")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(manager, w, r)
	})
	http.ListenAndServe(":8080", nil)
}

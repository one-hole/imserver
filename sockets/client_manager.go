package sockets

import (
	"fmt"
)

// ClientManager stands for the hub of clients
type ClientManager struct {
	clients    map[*Client]bool // 这里存储所有的连接信息
	register   chan *Client
	unregister chan *Client
	Broadcast  chan []byte
}

// NewManger returns an instance of ClientManger
func NewManger() *ClientManager {
	return &ClientManager{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

// Close release resources
func (manager *ClientManager) Close() {
	fmt.Println("release manager resources")
}

// Clients returns all connected clients
func (manager *ClientManager) Clients() map[*Client]bool {
	return manager.clients
}

// Exec || activate the manager
func (manager *ClientManager) Exec() {
	for {
		select {
		case client := <-manager.register:
			fmt.Println("Registing client ...")
			manager.clients[client] = true
		case client := <-manager.unregister:
			if _, ok := manager.clients[client]; ok {
				delete(manager.clients, client)
				client.destory()
			}
		case msgs := <-manager.Broadcast:
			for client := range manager.clients {
				select {
				case client.messages <- msgs:
				default:
					client.destory()
					delete(manager.clients, client)
				}
			}
		}
	}
}

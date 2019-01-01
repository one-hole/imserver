package sockets

import "fmt"

// ClientManager stands for the hub of clients
type ClientManager struct {
	clients    map[*Client]bool // 这里存储所有的连接信息
	register   chan *Client
	unregister chan *Client
}

// NewManger returns an instance of ClientManger
func NewManger() *ClientManager {
	return &ClientManager{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Close release resources
func (manager *ClientManager) Close() {
	fmt.Println("release manager resources")
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
		}
	}
}

package sockets

import "fmt"

// ClientManager stands for the hub of clients
type ClientManager struct {
	register   chan *Client
	unregister chan *Client
}

// NewManger returns an instance of ClientManger
func NewManger() *ClientManager {
	return &ClientManager{
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Close release resources
func (manager *ClientManager) Close() {
	fmt.Println("release manager resources")
}

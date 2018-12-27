package sockets

import "github.com/gorilla/websocket"

// Client 是存在于服务端对连接的抽象描述 & 每一个连接都需要初始化一个 Client Instance
type Client struct {
	conn     *websocket.Conn
	messages chan []byte
}

func (c *Client) readMessageFromClient() {

}

func (c *Client) writeMessageToClient() {

}

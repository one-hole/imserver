package sockets

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 5 * time.Second
	pongWait       = 6 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// Client 是存在于服务端对连接的抽象描述 & 每一个连接都需要初始化一个 Client Instance
type Client struct {
	manager  *ClientManager
	conn     *websocket.Conn
	messages chan []byte
}

func (c *Client) Conn() *websocket.Conn {
	return c.conn
}

// 这里暂时读消息只读心跳包
func (c *Client) readMessageFromClient() {
	defer func() {
		c.manager.unregister <- c
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait)) // 从 conn读取 最多等待 pongWait 的时间 & 这个语句用来第一次
	c.conn.SetPongHandler(func(string) error {       // 这里设置 Pong 消息的处理器 & 如果没有收到 Pong 消息 那就会读到 Error
		fmt.Println("Receive pong message...")
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// 这样的死循环一般来说开启一个 Goroutine
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
	}

}

func (c *Client) writeMessageToClient() {
	ticker := time.NewTicker(pingPeriod) // 每隔 pingPeriod 触发一次 Ping 操作
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case msg, ok := <-c.messages:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait)) // 写入 Conn 的时间
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)
			// n := len(c.messages)
			// for i := 0; i < n; i++ {
			// 	w.Write([]byte{'\n'})
			// 	w.Write(<-c.messages)
			// }
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			log.Println("Sending Ping message to client ...")
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Destory exports destory()
func (c *Client) Destory() {
	c.destory()
}

func (c *Client) destory() {
	delete(c.manager.clients, c)
	c.conn.Close()
	close(c.messages)
}

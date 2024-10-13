package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
}

func (c *Client) Read() {
	defer func() {
		c.hub.Unregister(c)
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(err.Error())
			}
			break
		}
		c.hub.Broadcast(msg, c)
	}
}

func (c *Client) Write(msg []byte) {
	if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
		return
	}
}

func (c *Client) Close() {
	defer c.conn.Close()
	msg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Closing connection")
	_ = c.conn.WriteMessage(websocket.CloseMessage, msg)
}

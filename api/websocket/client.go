package websocket

import (
	"fmt"
	"log"
	"os"

	"github.com/Amit152116Kumar/chess_server/movegen"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub   *Hub
	conn  *websocket.Conn
	color movegen.Color
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
		fmt.Fprintf(os.Stdout, "websocket msg from client color %s, --> %s", c.color.Name(), msg)
		c.hub.processRequestPacket(msg, c)
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

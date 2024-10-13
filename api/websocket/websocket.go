package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  32,
	WriteBufferSize: 32,
}
var wsConnections = WsConnections{}

func WsHandler(c *gin.Context) {
	id := c.Param("id")
	uid := uuid.MustParse(id)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, http.Header{
		"ws": []string{"ws"},
	})

	if err != nil {
		log.Println("WsHandler: ", err)
	}
	var hub *Hub
	hub, ok := wsConnections[uid]
	if !ok {
		hub = newHub()
		wsConnections[uid] = hub
	}
	client := &Client{
		hub:  hub,
		conn: conn,
	}
	hub.Register(client)
	go client.Read()

}

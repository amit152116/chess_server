package websocket

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/moveGen/models"
	"github.com/google/uuid"
)

type WsConnections map[uuid.UUID]*Hub

type Hub struct {
	clients    map[*Client]bool
	chessBoard *models.Board
}

func newHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) Register(client *Client) {
	n := len(h.clients)
	if n < 2 {
		h.clients[client] = true
	} else {
		client.Write([]byte("Room is full"))
		client.Close()
	}
}

func (h *Hub) Unregister(client *Client) {
	if _, ok := h.clients[client]; ok {
		client.Close()
		delete(h.clients, client)
	}
}

func (h *Hub) Broadcast(msg []byte, sender *Client) {
	fmt.Println(string(msg))
	if string(msg) == "start" {
		var board, _ = models.LoadFen(models.StartFEN)
		fmt.Println(board)
		h.chessBoard = board
	} else if string(msg) == "moves" {
		var movesList = h.chessBoard.GetAllMoves(false)
		var byteArr []byte
		fmt.Println(movesList)
		for from, moves := range movesList {
			byteArr = append(byteArr, byte(from))
			byteArr = append(byteArr, 0)
			byteArr = append(byteArr, byte(len(moves)))
			byteArr = append(byteArr, 0)
			for _, move := range moves {
				byteArr = append(byteArr, byte(move))
			}
		}
		sender.Write(byteArr)
	} else if len(msg) == 2 {
		h.chessBoard.UpdateBoard(int(msg[0]), int(msg[1]))
	} else {
		for client := range h.clients {
			if client != sender {
				client.Write(msg)
			}
		}
	}

}

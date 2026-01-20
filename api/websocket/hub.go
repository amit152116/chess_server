package websocket

import (
	"github.com/amit152116/chess_server/api/websocket/protocol"
	"github.com/amit152116/chess_server/movegen"
	"github.com/google/uuid"
)

type WsConnections map[uuid.UUID]*Hub

type Hub struct {
	clients    map[*Client]bool
	chessBoard *movegen.Board
}

func newHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) Register(client *Client) {
	n := len(h.clients)

	switch n {
	case 0:
		h.clients[client] = true
	case 1:
		h.clients[client] = true
		h.chessBoard, _ = movegen.LoadFen(movegen.StartFEN)
		h.Broadcast([]byte{1, 0})
	default:
		client.Write([]byte{1, 3})
	}
}

func (h *Hub) Unregister(client *Client) {
	if _, ok := h.clients[client]; ok {
		client.Close()
		delete(h.clients, client)
	}
}

// Broadcast Packet Structure (in bytes):
//
// HeaderPacket (4 bytes total):
// 1. 1 byte: Packet type
//   - Defines the type of the message (e.g., move, game state, error, etc.).
//
// 2. 2 bytes: Length of the body
//   - Specifies the length of the body to indicate how many bytes to expect for the message data.
//
// 3. 1 byte: Player identifier
//   - Identifies which player is sending the message (used to differentiate between the two players).
//
// Body (variable length):
// Contains the actual data based on the packet type.
//
// Packet Types:
//
// 1. Move Packet (0x01): Represents a chess move.
//   - Body:
//   - 1 byte: Starting square ('from' position, 0-63, representing the chessboard index).
//   - 1 byte: Ending square ('to' position, 0-63, representing the target square).
//   - 1 byte: Optional move metadata (e.g., promotion type or special move details).
//
// 2. Game State Packet (0x02): Represents the current state of the game.
//   - Body:
//   - 64 bytes: Board state (1 byte per square, each byte represents a piece or an empty square).
//   - 1 byte: Turn indicator (indicates which player's turn it is).
//   - 4 bytes: Time remaining for each player (2 bytes for each player, in seconds).
//
// 4. Chat Packet (0x04): Optional, used to send player-to-player chat messages.
//   - Body:
//   - 1 byte: Length of the chat message (indicates how many bytes the message contains).
//   - Variable: Chat message (in UTF-8 encoding).
//
// 5. Legal Moves Request Packet (0x05): Requests the available legal moves for the current player.
//   - Body:
//   - No additional data in the body (empty body).
//
// 6. Legal Moves Response Packet (0x06): Responds with the available legal moves for the current player.
//   - Body:
//   - 1 byte: Number of legal moves available (N).
//   - N bytes: List of legal moves (each move represented by a byte, where the byte indicates the target square).
//
// 7. Game Over Packet (0x07): Notifies players that the game has ended.
//   - Body:
//   - 1 byte: Result code (checkmate, stalemate, resignation, etc.)
//   - 1 byte: Winner identifier (indicates which player won; 0x00 for draw).
//   - Variable: Optional message or reason for the game ending.
//
// 9. Resignation Packet (0x09): Sent by a player to resign from the game.
//   - Body:
//   - 1 byte: Resigning player identifier.
//
// 10. Player Status Changed Packet (0x0A): Updates both players on the status of each player.
//   - Body:
//   - 1 byte: Player identifier (which player’s status is being updated).
//   - 1 byte: Status code (0x00 for online, 0x01 for offline).
//
// 11. Promotion Request Packet (0x0B): Notifies when a pawn is promoted.
//   - Body:
//   - 1 byte: Pawn’s position (0-63).
//   - 1 byte: New piece type (0x01 for queen, 0x02 for rook, 0x03 for bishop, 0x04 for knight).
//
// 12. Ping Packet (0x0C): A lightweight packet to check connectivity.
//   - Body:
//   - 0 bytes: No additional data.
//
// 13. Ping Response Packet (0x0D): Confirms server is alive.
//   - Body:
//   - 0 bytes: No additional data.
//
// 15. Undo Move Request Packet (0x0F): Requests to undo the last move.
//   - Body:
//   - 0 bytes: No additional data.
//
// 16. Undo Move Response Packet (0x10): Responds to the undo request.
//   - Body:
//   - 1 byte: Result code (success or failure).
//   - Variable: Optional message explaining the result.
//
// 17. Draw Offer Packet (0x11): A player offers a draw.
//   - Body:
//   - 1 byte: Player identifier (who is offering the draw).
//
// 18. Draw Response Packet (0x12): Response to a draw offer.
//   - Body:
//   - 1 byte: Result code (accept or reject).
//   - 1 byte: Player identifier (who accepted or rejected the draw).
func (h *Hub) Broadcast(responseBytes []byte) {
	for client := range h.clients {
		client.Write(responseBytes)
	}
}

func (h *Hub) SendToOpponent(responseBytes []byte, sender *Client) {
	for client := range h.clients {
		if client != sender {
			client.Write(responseBytes)
			break
		}
	}
}

func (h *Hub) processRequestPacket(msgPacket []byte, sender *Client) {
	packetType := protocol.RequestPacketType(msgPacket[0])
	data := msgPacket[1:]
	switch packetType {
	case protocol.Move:
		{
			go h.SendToOpponent(msgPacket, sender)
			packet := &protocol.MovePacket{}
			err := packet.Decode(data)
			if err != nil {
				return
			}
			h.chessBoard.UpdateBoard(packet.From, packet.To)
			responsePacket := &protocol.AcknowledgmentPacket{}
			responsePacket.Status = 0
			responsePacket.Message = "OK"
			responseBytes, err := responsePacket.Encode(sender.color)
			if err != nil {
				return
			}
			sender.Write(responseBytes)
		}
	case protocol.GameStateRequest:
		{
			responsePacket := &protocol.GameStatePacket{}
			responsePacket.Turn = h.chessBoard.Turn
			responsePacket.Board = []byte(h.chessBoard.GetFEN())
			responseBytes, err := responsePacket.Encode(sender.color)
			if err != nil {
				return
			}
			sender.Write(responseBytes)
		}
	case protocol.LegalMovesRequests:
		{
			responsePacket := &protocol.LegalMovePacket{}
			responsePacket.LegalMoves = h.chessBoard.GetAllMoves(false)
			responseBytes, err := responsePacket.Encode(sender.color)
			if err != nil {
				return
			}
			sender.Write(responseBytes)
		}
	case protocol.Promotion:
		{
			go h.SendToOpponent(msgPacket, sender)
			packet := &protocol.PromotionPacket{}
			err := packet.Decode(data)
			if err != nil {
				return
			}

			// todo update the board with the promotion piece
			responsePacket := &protocol.AcknowledgmentPacket{}
			responsePacket.Status = 0
			responsePacket.Message = "OK"
			responseBytes, err := responsePacket.Encode(sender.color)
			if err != nil {
				return
			}
			sender.Write(responseBytes)
		}
	case protocol.Resignation:
	case protocol.PingRequest:
	case protocol.UndoMoveRequest:
	case protocol.DrawOfferRequest:
	case protocol.AbortRequest:
	case protocol.Chat:
	default:
		sender.Write([]byte("Wrong Packet Code\n"))
	}
}

package protocol

import (
	"fmt"

	"github.com/Amit152116Kumar/chess_server/movegen"
)

type RequestPacket interface {
	Decode(data []byte) error
}

// MovePacket represents a move packet.
type MovePacket struct {
	header HeaderPacket
	From   byte
	To     byte
	Meta   string
}

func (packet *MovePacket) Decode(data []byte) error {
	err := packet.header.Decode(data[:4])
	if err != nil {
		return fmt.Errorf("failed to decode header: %w", err)
	}
	if len(data) != 4+packet.header.bodyLength { // 4 bytes for header + L bytes for From, To, Meta
		return fmt.Errorf("data too short, expected at least 7 bytes but got %d", len(data))
	}
	packet.From = data[4]
	packet.To = data[5]
	packet.Meta = string(data[6:])
	return nil
}

type ChatPacket struct {
	header  HeaderPacket
	Message string
	Length  int
}

func (packet *ChatPacket) Decode(data []byte) error {
	// TODO implement me
	panic("implement me")
}

type PromotionPacket struct {
	header         HeaderPacket
	Position       int
	PromotionPiece movegen.Piece
}

func (packet *PromotionPacket) Decode(data []byte) error {
	// TODO implement me
	panic("implement me")
}

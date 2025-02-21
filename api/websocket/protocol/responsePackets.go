package protocol

import (
	"bytes"
	"fmt"
	"time"

	"github.com/Amit152116Kumar/chess_server/movegen"
	"github.com/Amit152116Kumar/chess_server/utils"
)

type ResponsePacket interface {
	Encode(color movegen.Color) ([]byte, error)
}

// AcknowledgmentPacket represents the successful processing of a request
type AcknowledgmentPacket struct {
	header  HeaderPacket
	Status  byte
	Message string
}

func (packet *AcknowledgmentPacket) Encode(color movegen.Color) ([]byte, error) {
	buffer := new(bytes.Buffer)
	buffer.WriteByte(packet.Status)
	buffer.WriteByte(byte(len(packet.Message)))
	buffer.Write([]byte(packet.Message))

	packet.header.packetType = PacketType(GameStateResponse)
	packet.header.Player = color
	packet.header.bodyLength = buffer.Len()
	headerBytes, err := packet.header.Encode(color)
	if err != nil {
		return nil, fmt.Errorf("failed to encode header: %w", err)
	}

	return append(headerBytes, buffer.Bytes()...), nil
}

// GameStatePacket represents a game state packet.
type GameStatePacket struct {
	header HeaderPacket
	Board  []byte
	Turn   movegen.Color
	Time   time.Duration
}

func (packet *GameStatePacket) Encode(color movegen.Color) ([]byte, error) {
	// TODO structure the packet more correctly
	buffer := new(bytes.Buffer)
	buffer.WriteByte(byte(len(packet.Board)))
	buffer.Write(packet.Board)
	buffer.WriteByte(byte(packet.Turn))

	packet.header.packetType = PacketType(GameStateResponse)
	packet.header.Player = color
	packet.header.bodyLength = buffer.Len()
	headerBytes, err := packet.header.Encode(color)
	if err != nil {
		return nil, fmt.Errorf("failed to encode header: %w", err)
	}

	return append(headerBytes, buffer.Bytes()...), nil
}

// LegalMovePacket represent the available legalMoves for current gameState
type LegalMovePacket struct {
	header     HeaderPacket
	LegalMoves map[int][]int
}

func (packet *LegalMovePacket) Encode(color movegen.Color) ([]byte, error) {
	buffer := new(bytes.Buffer)
	buffer.WriteByte(byte(len(packet.LegalMoves)))
	for piecePos, moves := range packet.LegalMoves {
		// 1 byte: Piece position
		buffer.WriteByte(byte(piecePos))

		// 1 byte: Number of legal moves
		buffer.WriteByte(byte(len(moves)))

		// N bytes: Legal move positions
		for _, move := range moves {
			buffer.WriteByte(byte(move))
		}
	}

	packet.header.packetType = PacketType(LegalMovesResponse)
	packet.header.Player = color
	packet.header.bodyLength = buffer.Len()
	headerBytes, err := packet.header.Encode(color)
	if err != nil {
		return nil, fmt.Errorf("failed to encode header: %w", err)
	}

	return append(headerBytes, buffer.Bytes()...), nil
}

type GameOverPacket struct {
	header     HeaderPacket
	GameResult utils.GameResult
	Winner     movegen.Color
	Message    string
}

func (packet *GameOverPacket) Encode(color movegen.Color) ([]byte, error) {
	buffer := new(bytes.Buffer)

	packet.header.packetType = PacketType(GameStateResponse)
	packet.header.Player = color
	packet.header.bodyLength = buffer.Len()
	headerBytes, err := packet.header.Encode(color)
	if err != nil {
		return nil, fmt.Errorf("failed to encode header: %w", err)
	}

	return append(headerBytes, buffer.Bytes()...), nil
}

type PlayerStatusPacket struct {
	header   HeaderPacket
	PlayerId byte
	Status   bool
}

func (packet *PlayerStatusPacket) Encode(color movegen.Color) ([]byte, error) {
	buffer := new(bytes.Buffer)

	packet.header.packetType = PacketType(GameStateResponse)
	packet.header.Player = color
	packet.header.bodyLength = buffer.Len()
	headerBytes, err := packet.header.Encode(color)
	if err != nil {
		return nil, fmt.Errorf("failed to encode header: %w", err)
	}

	return append(headerBytes, buffer.Bytes()...), nil
}

type UndoMovePacket struct {
	header  HeaderPacket
	Result  bool
	Message string
}

func (packet *UndoMovePacket) Encode(color movegen.Color) ([]byte, error) {
	buffer := new(bytes.Buffer)

	packet.header.packetType = PacketType(GameStateResponse)
	packet.header.Player = color
	packet.header.bodyLength = buffer.Len()
	headerBytes, err := packet.header.Encode(color)
	if err != nil {
		return nil, fmt.Errorf("failed to encode header: %w", err)
	}

	return append(headerBytes, buffer.Bytes()...), nil
}

type DrawResponsePacket struct {
	header  HeaderPacket
	Result  bool
	Message string
}

func (packet *DrawResponsePacket) Encode(color movegen.Color) ([]byte, error) {
	buffer := new(bytes.Buffer)

	packet.header.packetType = PacketType(GameStateResponse)
	packet.header.Player = color
	packet.header.bodyLength = buffer.Len()
	headerBytes, err := packet.header.Encode(color)
	if err != nil {
		return nil, fmt.Errorf("failed to encode header: %w", err)
	}

	return append(headerBytes, buffer.Bytes()...), nil
}

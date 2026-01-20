package protocol

import (
	"bytes"
	"encoding/binary"

	"github.com/amit152116/chess_server/movegen"
	"github.com/amit152116/chess_server/myErrors"
)

type HeaderPacket struct {
	packetType PacketType
	bodyLength int
	Player     movegen.Color
}

func (h *HeaderPacket) Encode(color movegen.Color) ([]byte, error) {
	h.Player = color
	buffer := new(bytes.Buffer)
	if err := buffer.WriteByte(byte(h.packetType)); err != nil {
		return nil, err
	}
	if h.bodyLength > 0xFFFF {
		return nil, myErrors.ErrBodyLenTooLarge
	}
	lenByte := []byte{byte(h.bodyLength >> 8), byte(h.bodyLength & 0xFF)}
	if _, err := buffer.Write(lenByte); err != nil {
		return nil, err
	}
	if err := buffer.WriteByte(byte(h.Player)); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (h *HeaderPacket) Decode(data []byte) error {
	h.packetType = PacketType(data[0])
	h.bodyLength = int(binary.BigEndian.Uint16(data[1:3]))
	h.Player = movegen.Color(data[3])
	return nil
}

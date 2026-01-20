package protocol

import (
	"fmt"
	"testing"

	"github.com/amit152116/chess_server/movegen"
	"github.com/go-playground/assert/v2"
)

func TestHeaderPacket_Encode(t *testing.T) {
	header := &HeaderPacket{}

	header.packetType = PacketType(Acknowledgement)
	header.Player = movegen.White
	header.bodyLength = 65535

	bytes, err := header.Encode(0)
	fmt.Println(bytes, err)
	if err == nil {
		headerCopy := header
		assert.IsEqual(header, headerCopy.Decode(bytes))
	}
}

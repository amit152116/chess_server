package protocol

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/movegen"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestHeaderPacket_Encode(t *testing.T) {
	var header = &HeaderPacket{}

	header.packetType = PacketType(Acknowledgement)
	header.Player = movegen.White
	header.bodyLength = 65535

	var bytes, err = header.Encode(0)
	fmt.Println(bytes, err)
	if err == nil {
		var headerCopy = header
		assert.IsEqual(header, headerCopy.Decode(bytes))
	}

}

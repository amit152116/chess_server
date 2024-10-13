package moveGen

import (
	"github.com/Amit152116Kumar/chess_server/moveGen/models"
)

func Generate() {
	var board, err = models.LoadFen(models.Chess960FEN[2])
	if err != nil {
		println(err.Error())
		return
	}
	println(len(board.Bitboard))
}

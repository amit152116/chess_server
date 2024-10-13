package models

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestBoard_GetFEN(t *testing.T) {
	var board, err = LoadFen(StartFEN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		assert.IsEqual(board.GetFEN(), StartFEN)
	}
}

func TestBoard_WhichPieceExists(t *testing.T) {
	var board, _ = LoadFen(StartFEN)
	for i := 0; i < 64; i++ {
		var pieceName, color = GetNameAndColor(board.WhichPieceExists(i))
		fmt.Println(i, "->", pieceName.Name(), color.Name())
	}
}
func TestBoard_Copy(t *testing.T) {

}

func TestBoard_Copy2(t *testing.T) {

}

func TestBoard_IsKingInCheck(t *testing.T) {

}

func TestBoard_GetLegalMoves(t *testing.T) {

}
func TestBoard_GetMoves(t *testing.T) {

}

func TestBoard_GetAllMoves(t *testing.T) {
	var board, err = LoadFen(StartFEN)
	if err != nil {
		fmt.Println(err)
	}
	var moves = board.GetAllMoves(false)
	fmt.Println(moves)
	board.UpdateBoard(15, 31)
	board.DrawBoard(nil)

	moves = board.GetAllMoves(false)
	fmt.Println(moves)
	board.UpdateBoard(53, 37)
	board.DrawBoard(nil)
	
	moves = board.GetAllMoves(false)
	fmt.Println(moves)
}

func TestBoard_DrawBoard(t *testing.T) {
	var board, _ = LoadFen(StartFEN)
	board.DrawBoard(nil)
	t.Run("Get King Position", func(t *testing.T) {
		var pos = board.getKingPosition(White)
		fmt.Println("white king Pos: ", pos)
		pos = board.getKingPosition(Black)
		fmt.Println("black king pos: ", pos)
	})
}

package moveGen

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/moveGen/models"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestMoveGen(t *testing.T) {
	var board, err = models.LoadFen(models.StartFEN)
	if err != nil {
		println(err.Error())
		return
	}
	t.Run("Load Fen", func(t *testing.T) {

		assert.Equal(t, models.StartFEN, board.GetFEN())
	})

	t.Run("Get SquareIDX", func(t *testing.T) {
		var idx, err = models.GetSquarePos("h8")
		if err != nil {
			println(err.Error())
		}
		println("Square idx : ", idx)
	})

	t.Run("Get Square Name", func(t *testing.T) {
		println(models.GetSquareName(32))
	})

	t.Run("PieceName Exists", func(t *testing.T) {

		for i := 0; i < 64; i++ {
			var _, color = models.GetNameAndColor(board.WhichPieceExists(i))
			fmt.Println(i, color)
		}
	})

	//t.Run("Knight Move", func(t *testing.T) {
	//	var moves = board.knightMoves(board, 32)
	//
	//	board.DrawBoard(moves)
	//})
	//
	//t.Run("Bishop Move", func(t *testing.T) {
	//	var moves = bishopMoves(board, 33)
	//
	//	board.DrawBoard(moves)
	//})
	//
	//t.Run("Rook Moves", func(t *testing.T) {
	//	var moves = rookMoves(board, 35)
	//	board.DrawBoard(moves)
	//})
	//t.Run("Queen Moves", func(t *testing.T) {
	//	var moves = queenMoves(board, 35)
	//	board.DrawBoard(moves)
	//})
	//
	//t.Run("Pawn Forward", func(t *testing.T) {
	//	var moves = pawnForwardMoves(board, 55)
	//	board.DrawBoard(moves)
	//})
	//
	//t.Run("Making Move", func(t *testing.T) {
	//	var moves []int
	//	board.DrawBoard(moves)
	//	updateBoard(board, 9, 25)
	//
	//	println()
	//	println()
	//	board.DrawBoard(moves)
	//})

	t.Run("Get All Position", func(t *testing.T) {
		board.GetAllPosition()
	})

	t.Run("Get All Moves", func(t *testing.T) {
		var moves = board.GetAllMoves(false)

		for pieceName, innerMap := range moves {
			fmt.Printf("%d \t-->\n", pieceName)
			for idx, moveIndices := range innerMap {
				fmt.Printf("\t %d \t -> \t ", idx)
				fmt.Print(moveIndices)
				fmt.Println()
			}
			fmt.Println()
		}
	})

	t.Run("Get Move list", func(t *testing.T) {
		var moves = board.GetMoves(models.Knight, 6, false)
		fmt.Println(moves)
		board.DrawBoard(moves)
	})

	t.Run("Is king in Check", func(t *testing.T) {
		board.IsKingInCheck(models.White)
		board.IsKingInCheck(models.Black)
	})

}

func BenchmarkGenerate(b *testing.B) {
	var board, _ = models.LoadFen(models.StartFEN)
	board.Copy()
}

func BenchmarkGenerate2(b *testing.B) {
	var board, _ = models.LoadFen(models.StartFEN)
	board.Copy2()
}

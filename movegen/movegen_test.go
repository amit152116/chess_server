package movegen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMoveGen(t *testing.T) {
	board, err := LoadFen(StartFEN)
	if err != nil {
		println(err.Error())
		return
	}
	t.Run("Load Fen", func(t *testing.T) {
		assert.Equal(t, StartFEN, board.GetFEN())
	})

	t.Run("Get Square Name", func(t *testing.T) {
		println(getSquareName(32))
	})

	t.Run("Piece Exists", func(t *testing.T) {
		for i := 0; i < 64; i++ {
			_, color := getNameAndColor(board.WhichPieceExists(byte(i)))
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

	t.Run("Get All Moves", func(t *testing.T) {
		moves := board.GetAllMoves(false)

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
		moves := board.GetPartialMoves(Knight, 6, false)
		fmt.Println(moves)
		board.DrawBoard(moves)
	})

	t.Run("Is king in Check", func(t *testing.T) {
		board.IsKingInCheck(White)
		board.IsKingInCheck(Black)
	})
}

func BenchmarkGenerate(b *testing.B) {
	board, _ := LoadFen(StartFEN)
	board.Copy()
}

func TestBoard_GetFEN(t *testing.T) {
	board, err := LoadFen(StartFEN)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		assert.IsEqual(board.GetFEN(), StartFEN)
	}
}

func TestBoard_WhichPieceExists(t *testing.T) {
	board, _ := LoadFen(StartFEN)
	for i := 0; i < 64; i++ {
		pieceName, color := getNameAndColor(board.WhichPieceExists(byte(i)))
		fmt.Println(i, "->", pieceName.Name(), color.Name())
	}
}

func TestBoard_GetAllMoves(t *testing.T) {
	board, err := LoadFen(StartFEN)
	if err != nil {
		fmt.Println(err)
	}
	moves := board.GetAllMoves(false)
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
	board, _ := LoadFen(StartFEN)
	board.DrawBoard(nil)
	t.Run("Get King Position", func(t *testing.T) {
		pos := board.getKingPosition(White)
		fmt.Println("white king Pos: ", pos)
		pos = board.getKingPosition(Black)
		fmt.Println("black king pos: ", pos)
	})
}

func TestValidateFEN(t *testing.T) {
	for idx, fen := range Chess960FEN {
		rules := strings.Split(fen, " ")
		fmt.Println(idx, " --> ", validateBoard(rules[0]), validateTurnToMove(rules[1]), validateCastlingRule(rules[2]), validateEnpassant(rules[3]), validateHalfMove(rules[4]), validateFullMove(rules[5]))
	}
}

// func TestValidate_CastlingRule(t *testing.T) {
// 	rule := "KQkq"
// 	res := ""
// 	var arr []string
// 	reverse := func(res string,) {
// 	}
//
// 	reverse()
// }

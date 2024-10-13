package models

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var StartFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Board struct {
	Bitboard        map[uint8]uint64
	EnPassantSquare int
	Turn            Color
	HalfMove        int
	FullMove        int
	Castle          uint8
	Score           int
}

var invalidFen = errors.New("invalid FEN")

func LoadFen(fen string) (*Board, error) {
	if len(fen) == 0 {
		fen = StartFEN
	}
	var board = &Board{}
	var rules = strings.Split(fen, " ")

	board.Bitboard = getBitBoard(rules[0])
	if rules[1] == "w" {
		board.Turn = White
	} else if rules[1] == "b" {
		board.Turn = Black
	} else {
		return nil, invalidFen
	}
	var err error = nil
	board.Castle, err = getCastlingRule(rules[2])
	if err != nil {
		return nil, err
	}
	board.EnPassantSquare, err = enPassantRule(rules[3])
	if err != nil {
		return nil, err
	}
	board.HalfMove, _ = strconv.Atoi(rules[4])
	board.FullMove, _ = strconv.Atoi(rules[5])

	return board, nil
}

func getCastlingRule(rule string) (uint8, error) {
	if len(rule) > 4 {
		return 0, invalidFen
	}
	var castle uint8 = 0
	if strings.Contains(rule, "K") {
		castle |= 1 << 0
	}
	if strings.Contains(rule, "Q") {
		castle |= 1 << 1
	}
	if strings.Contains(rule, "k") {
		castle |= 1 << 2
	}
	if strings.Contains(rule, "q") {
		castle |= 1 << 3
	}

	return castle, nil
}

func enPassantRule(rule string) (int, error) {
	if rule == "-" {
		return -1, nil
	} else {
		return GetSquarePos(rule)
	}
}

func getBitBoard(board string) map[uint8]uint64 {
	var bitboard = make(map[uint8]uint64)
	var rank, file = 7, 0

	for _, char := range board {
		if char == '/' {
			file = 0
			rank--
		} else if char >= '0' && char <= '9' {
			emptyRow := int(char - '0')
			file += emptyRow
		} else {
			pieceName := fromNotation(strings.ToLower(string(char)))

			if char >= 'A' && char <= 'Z' {
				bitboard[GetPiece(pieceName, White)] |= 1 << (rank*8 + file)
			} else {

				bitboard[GetPiece(pieceName, Black)] |= 1 << (rank*8 + file)
			}

			file++
		}
	}
	return bitboard
}

func (b *Board) GetFEN() string {
	var fen string

	for rank := 7; rank >= 0; rank-- {
		var emptySqu = 0

		for file := 0; file < 8; file++ {
			var pos = rank*8 + file
			var pieceName, color = GetNameAndColor(b.WhichPieceExists(pos))
			if pieceName == 0 {
				emptySqu++
				continue
			}
			if emptySqu > 0 {
				fen += fmt.Sprintf("%d", emptySqu)
			}

			var notation = pieceName.Notation()
			if color == White {
				notation = strings.ToUpper(notation)
			}
			fen += notation
		}
		if emptySqu > 0 {
			fen += fmt.Sprintf("%d", emptySqu)
		}
		if rank > 0 {
			fen += "/"
		}
	}
	fen += " "
	if b.Turn == White {
		fen += "w"
	} else {
		fen += "b"
	}
	fen += " "
	if b.Castle&(1<<0) != 0 {
		fen += "K"
	}
	if b.Castle&(1<<1) != 0 {
		fen += "Q"
	}
	if b.Castle&(1<<2) != 0 {
		fen += "k"
	}
	if b.Castle&(1<<3) != 0 {
		fen += "q"
	}
	fen += " "
	fen += GetSquareName(b.EnPassantSquare)
	fen += " "
	fen += fmt.Sprintf("%d", b.HalfMove)
	fen += " "
	fen += fmt.Sprintf("%d", b.FullMove)

	return fen
}

func (b *Board) WhichPieceExists(pos int) uint8 {
	// todo there is coming negative pos test it and solve it
	for piece, positions := range b.Bitboard {

		if positions&(1<<pos) != 0 {
			return piece
		}
	}
	return 0
}

func (b *Board) IsPieceExists(pieceName PieceName, color Color, pos int) bool {
	if b.Bitboard[GetPiece(pieceName, color)]>>pos&1 != 0 {
		return true
	}
	return false
}

func (b *Board) DrawBoard(moves []int) {

	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			var pos = rank*8 + file
			var piece = b.WhichPieceExists(pos)
			if slices.Contains(moves, pos) {
				fmt.Printf("*\t")
				continue
			}
			fmt.Printf("%d\t", piece)

		}
		fmt.Println()
	}

}

func (b *Board) GetAllMoves(attackOnly bool) map[int][]int {
	var moves = make(map[int][]int)
	for piece := range b.Bitboard {
		var pieceName, color = GetNameAndColor(piece)
		if color != b.Turn {
			continue
		}

		for i := 0; i < 64; i++ {
			var moveList []int
			if attackOnly {
				moveList = b.GetMoves(pieceName, i, attackOnly)
			} else {
				moveList = b.GetLegalMoves(pieceName, i)
			}
			//todo there is some problem here, solve it later
			if len(moveList) != 0 {
				moves[i] = moveList
			}
		}
	}
	return moves
}

func (b *Board) GetLegalMoves(pieceName PieceName, pos int) []int {
	var moves = b.GetMoves(pieceName, pos, false)

	var legalMoves []int
	for _, targetPos := range moves {
		var tempBoard = b.Copy()
		var targetPiece = tempBoard.WhichPieceExists(targetPos)
		tempBoard.setPiece(GetPiece(pieceName, tempBoard.Turn), targetPiece, pos, targetPos)
		if !tempBoard.IsKingInCheck(b.Turn) {
			legalMoves = append(legalMoves, targetPos)
		}
	}
	return legalMoves
}

func (b *Board) GetMoves(pieceName PieceName, pos int, attackMoves bool) []int {
	var moves []int
	if !b.IsPieceExists(pieceName, b.Turn, pos) {
		return moves
	}
	switch pieceName {
	case Pawn:
		moves = b.pawnAttackMoves(pos)
		if !attackMoves {
			moves = append(moves, b.pawnMoves(pos)...)
		}
	case Knight:
		moves = b.knightMoves(pos)
	case Bishop:
		moves = b.bishopMoves(pos)
	case Rook:
		moves = b.rookMoves(pos)
	case Queen:
		moves = b.queenMoves(pos)
	case King:
		moves = b.kingMoves(pos)
		if !attackMoves {
			moves = append(moves, b.castlingKingSide(pos)...)
			moves = append(moves, b.castlingQueenSide(pos)...)
		}
	}
	return moves
}

func (b *Board) Copy() *Board {
	var newBoard = &Board{
		Bitboard:        b.Bitboard,
		EnPassantSquare: b.EnPassantSquare,
		Turn:            b.Turn,
		HalfMove:        b.HalfMove,
		FullMove:        b.FullMove,
		Castle:          b.Castle,
		Score:           b.Score,
	}
	newBoard.Bitboard = make(map[uint8]uint64)

	for piece, positions := range b.Bitboard {
		newBoard.Bitboard[piece] = positions
	}

	return newBoard
}

func (b *Board) IsKingInCheck(color Color) bool {
	var KingPos = b.getKingPosition(color)
	//todo change getallmoves to other function
	var moveList = b.GetAllMoves(true)

	for _, moves := range moveList {
		if slices.Contains(moves, KingPos) {
			return true
		}
	}
	return false
}

func (b *Board) getKingPosition(color Color) int {
	var kingPos = b.Bitboard[GetPiece(King, color)]
	var pos = 0
	for ; kingPos != 1; pos++ {
		kingPos = kingPos >> 1
	}
	return pos
}

func (b *Board) knightMoves(pos int) []int {
	var moves []int
	var offset = []int{-17, -15, -10, -6, 6, 10, 15, 17}

	var currRank = GetRank(pos)
	var currFile = GetFile(pos)
	for i := 0; i < len(offset); i++ {
		var newPos = pos + offset[i]
		if !isOnBoard(newPos) {
			continue
		}

		newFile := GetFile(newPos)
		newRank := GetRank(newPos)

		if abs(currRank-newRank) > 2 || abs(currFile-newFile) > 2 {
			continue
		}
		var _, color = GetNameAndColor(b.WhichPieceExists(newPos))
		if color != b.Turn {
			moves = append(moves, newPos)
		}
	}
	return moves
}

func (b *Board) bishopMoves(pos int) []int {
	var moves []int
	var offset = []int{-9, -7, 7, 9}

	for i := 0; i < len(offset); i++ {
		var newPos = pos + offset[i]
		var oldPos = pos
		for isOnBoard(newPos) && GetRank(newPos) != GetRank(oldPos) {
			var _, color = GetNameAndColor(b.WhichPieceExists(newPos))
			if color == b.Turn {
				break
			}
			moves = append(moves, newPos)
			if color != 0 {
				break
			}
			oldPos = newPos
			newPos += offset[i]
		}
	}
	return moves
}

func (b *Board) rookMoves(pos int) []int {
	var moves []int
	var offset = []int{-1, 1, -8, 8}

	for i := 0; i < len(offset); i++ {
		var newPos = pos + offset[i]
		var oldPos = pos
		for isOnBoard(newPos) && (GetRank(oldPos) == GetRank(newPos) ||
			GetFile(oldPos) == GetFile(newPos)) {
			var _, color = GetNameAndColor(b.WhichPieceExists(newPos))

			if color == b.Turn {
				break

			}
			moves = append(moves, newPos)
			if color != 0 {
				break
			}
			oldPos = newPos
			newPos += offset[i]
		}
	}
	return moves
}

func (b *Board) queenMoves(pos int) []int {
	var moves []int
	moves = append(moves, b.rookMoves(pos)...)
	moves = append(moves, b.bishopMoves(pos)...)

	return moves
}

func (b *Board) pawnAttackMoves(pos int) []int {
	var moves []int
	var offset = []int{7, 9}

	for i := 0; i < len(offset); i++ {
		var newPos = pos + offset[i]*b.Turn.Value()
		if !isOnBoard(newPos) {
			continue
		}
		var _, color = GetNameAndColor(b.WhichPieceExists(newPos))
		if color != b.Turn && color != 0 {
			moves = append(moves, newPos)
		}
	}
	return moves
}

func (b *Board) pawnMoves(pos int) []int {
	var moves []int
	var offset = 8 * b.Turn.Value()
	var newPos = pos + offset
	if isOnBoard(newPos) {
		var pieceName, _ = GetNameAndColor(b.WhichPieceExists(newPos))
		if pieceName == 0 {
			moves = append(moves, newPos)
			var color = b.Turn
			var currentRank = GetRank(pos)

			if (color == White && currentRank == 1) || (color == Black && currentRank == 6) {
				newPos += offset
				var pieceName, _ = GetNameAndColor(b.WhichPieceExists(newPos))
				if pieceName == 0 {
					moves = append(moves, newPos)
				}
			}
		}
	}

	//en-passant move
	if b.EnPassantSquare != -1 {
		if GetRank(pos) == GetRank(b.EnPassantSquare) &&
			abs(GetFile(pos)-GetFile(b.EnPassantSquare)) == 1 {
			moves = append(moves, b.EnPassantSquare+offset)
		}
	}
	return moves
}

func (b *Board) kingMoves(pos int) []int {
	var moves []int
	var offset = []int{-9, -8, -7, -1, 1, 7, 8, 9}

	for i := 0; i < len(offset); i++ {
		var newPos = pos + offset[i]
		if !isOnBoard(newPos) {
			continue
		}
		var _, color = GetNameAndColor(b.WhichPieceExists(newPos))
		if color == b.Turn {
			continue
		}
		moves = append(moves, newPos)

	}
	return moves
}

func (b *Board) castlingQueenSide(pos int) []int {
	var moves []int
	var bitshift int
	if b.Turn == White {
		bitshift = 1
	} else {
		bitshift = 3
	}
	var isPossible = b.Castle >> bitshift & 1
	//todo check it later
	if isPossible != 0 {
		var rookPos = pos - 4
		var pathClear = true

		for i := rookPos + 1; i < pos; i++ {
			var pieceName, _ = GetNameAndColor(b.WhichPieceExists(i))
			if pieceName != 0 {
				pathClear = false
				break
			}
		}
		if pathClear {
			moves = append(moves, pos-2)
		}
		println()
	}
	return moves
}
func (b *Board) castlingKingSide(pos int) []int {
	var moves []int
	var bitshift int
	if b.Turn == White {
		bitshift = 0
	} else {
		bitshift = 2
	}
	//todo check it later
	var isPossible = b.Castle >> bitshift & 1
	if isPossible != 0 {
		var rookPos = pos + 3
		var pathClear = true

		for i := pos + 1; i < rookPos; i++ {
			var pieceName, _ = GetNameAndColor(b.WhichPieceExists(i))
			if pieceName != 0 {
				pathClear = false
				break
			}
		}
		if pathClear {
			moves = append(moves, pos-2)
		}
		println()
	}
	return moves
}
func (b *Board) UpdateBoard(from int, to int) {
	var pieceName, color = GetNameAndColor(b.WhichPieceExists(from))
	var targetPieceName, targetColor = GetNameAndColor(b.WhichPieceExists(to))

	b.EnPassantSquare = -1
	b.HalfMove++
	if targetPieceName != 0 {
		b.HalfMove = 0
		b.Score -= targetPieceName.Value() * targetColor.Value()
	}

	switch pieceName {
	case Pawn:
		b.HalfMove = 0
		//Update en-passant squarePos
		if abs(from-to) == 16 {
			b.EnPassantSquare = to
		}
		//Make en-passant Move
		if b.EnPassantSquare != -1 && abs(to-b.EnPassantSquare) == 8 {
			b.Score -= targetColor.Value()
		}
		if b.isPawnPromoted(to) {
			var promotionPiece = Queen
			b.Score += promotionPiece.Value() * b.Turn.Value()
		}
	case Rook:
		// setting casting bit to zero : binary(3) -> 11 && binary(12) -> 1100
		if b.Turn == White {
			b.Castle = b.Castle & ^uint8(3)
		} else {
			b.Castle = b.Castle & ^uint8(12)
		}
	case King:
		if b.Turn == White {
			b.Castle = b.Castle & ^uint8(3)
		} else {
			b.Castle = b.Castle & ^uint8(12)
		}
		// If castling move than also move the rook
		if abs(from-to) > 1 {
			var currRook int
			var targetRook int
			if from > to {
				currRook = to - 2
				targetRook = to + 1
			} else {
				currRook = to + 1
				targetRook = to - 1
			}
			b.setPiece(GetPiece(Rook, b.Turn), GetPiece(0, 0), currRook, targetRook)
		}
	default:
	}
	b.setPiece(GetPiece(pieceName, color), GetPiece(targetPieceName, targetColor), from, to)
	b.rotateTurn()
}

func (b *Board) isPawnPromoted(pos int) bool {
	if (b.Turn == White && pos > 55) || (b.Turn == Black && pos < 8) {
		return true
	}
	return false
}
func (b *Board) rotateTurn() {
	if b.Turn == White {
		b.Turn = Black
	} else {
		b.FullMove++
		b.Turn = White
	}
}

func (b *Board) setPiece(piece, targetPiece uint8, from, to int) {
	//todo add nor operation to solve all cases
	if targetPiece != 0 {
		b.Bitboard[targetPiece] &= ^(1 << to)
	}

	b.Bitboard[piece] ^= 1 << from
	b.Bitboard[piece] ^= 1 << to
}

func GetSquarePos(square string) (int, error) {
	var file = square[0] - 97
	var rank = square[1] - '1'
	var pos = rank*8 + file
	if pos >= 64 || pos < 0 {
		return 0, invalidFen
	}
	return int(pos), nil
}

func GetSquareName(pos int) string {
	if pos >= 64 || pos < 0 {
		return "-"
	}
	var file = GetFile(int(pos))
	var rank = GetRank(int(pos))
	return fmt.Sprintf("%c%d", rune(file+97), rank+1)
}

func GetRank(pos int) int {
	return pos >> 3
}

func GetFile(pos int) int {
	return pos & 7
}

func isOnBoard(targetPos int) bool {
	if targetPos >= 64 || targetPos < 0 {
		return false
	}
	return true

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (b *Board) GetAllPosition() {
	for piece, positions := range b.Bitboard {
		var pieceName, color = GetNameAndColor(piece)
		fmt.Printf("%s %s   \t -> \t", color.Name(), pieceName.Name())

		for i := 0; i < 64; i++ {
			if positions>>i&1 != 0 {
				fmt.Printf("%d, ", i)
			}
		}
		fmt.Println()
	}

}

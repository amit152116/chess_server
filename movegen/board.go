package movegen

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/Amit152116Kumar/chess_server/myErrors"
)

type Board struct {
	Bitboard        map[byte]uint64
	EnPassantSquare byte
	Turn            Color
	HalfMove        int
	FullMove        int
	Castle          byte
	Score           int
}

func LoadFen(fen string) (*Board, error) {
	rules := strings.Split(fen, " ")
	if !ValidateFEN(rules) {
		return nil, myErrors.InvalidFen
	}
	return parseFEN(rules), nil
}

func parseFEN(rules []string) *Board {
	board := &Board{}

	switch rules[1] {
	case "w":
		board.Turn = White
	case "b":
		board.Turn = Black
	}
	board.Bitboard = getBitBoard(rules[0])
	board.Castle = getCastlingRule(rules[2])
	board.EnPassantSquare = enPassantRule(rules[3])
	board.HalfMove, _ = strconv.Atoi(rules[4])
	board.FullMove, _ = strconv.Atoi(rules[5])

	return board
}

func getCastlingRule(rule string) byte {
	var castle byte = 0

	if rule == "-" {
		return castle
	}

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

	return castle
}

func enPassantRule(rule string) byte {
	if rule == "-" {
		return 255
	}
	file := rule[0] - 'a'
	rank := rule[1] - '1'
	pos := rank*8 + file
	return pos
}

func getBitBoard(board string) map[byte]uint64 {
	bitboard := make(map[byte]uint64)
	rank, file := 7, 0

	for _, char := range board {
		if char == '/' {
			file = 0
			rank--
		} else if unicode.IsDigit(char) {
			emptyRow := int(char - '0')
			file += emptyRow
		} else {
			pieceName := fromNotation(strings.ToLower(string(char)))

			if unicode.IsUpper(char) {
				bitboard[getPiece(pieceName, White)] |= 1 << (rank*8 + file)
			} else {
				bitboard[getPiece(pieceName, Black)] |= 1 << (rank*8 + file)
			}
			file++
		}
	}
	return bitboard
}

func (b *Board) GetFEN() string {
	var fen string

	for rank := 7; rank >= 0; rank-- {
		emptySqu := 0

		for file := 0; file < 8; file++ {
			pos := rank*8 + file
			pieceName, color := getNameAndColor(b.WhichPieceExists(byte(pos)))
			if pieceName == 0 {
				emptySqu++
				continue
			}
			if emptySqu > 0 {
				fen += fmt.Sprintf("%d", emptySqu)
			}

			notation := pieceName.Notation()
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
	fen += getSquareName(b.EnPassantSquare)
	fen += " "
	fen += fmt.Sprintf("%d", b.HalfMove)
	fen += " "
	fen += fmt.Sprintf("%d", b.FullMove)

	return fen
}

// GetAllMovesBytes Packet Structure for Move List:
//
// The packet represents the moves available for each piece on a chessboard.
// Data is serialized into a byte array and sent via WebSocket.
//
// Packet Format:
//  1. No. of pieces in the moveList
//     - For each piece (represented by "from"):
//     a. byte(from): The position of the piece (starting position).
//     b. byte(len(moves)): The number of possible moves from the current position.
//     c. For each move:
//     - byte(move): The destination position (move).
func (b *Board) GetAllMovesBytes() []byte {
	var byteArr []byte
	moveList := b.GetAllMoves(false)

	byteArr = append(byteArr, byte(len(moveList)))
	for from, moves := range moveList {
		byteArr = append(byteArr, byte(from))
		byteArr = append(byteArr, byte(len(moves)))
		for _, move := range moves {
			byteArr = append(byteArr, byte(move))
		}
	}

	return byteArr
}

func (b *Board) GetAllMoves(attackOnly bool) map[int][]int {
	moves := make(map[int][]int)
	for piece := range b.Bitboard {
		pieceName, color := getNameAndColor(piece)
		if color != b.Turn {
			continue
		}

		for i := 0; i < 64; i++ {
			var moveList []int
			if attackOnly {
				moveList = b.GetPartialMoves(pieceName, i, attackOnly)
			} else {
				moveList = b.GetLegalMoves(pieceName, i)
			}
			// TODO there is some problem here, solve it later
			if len(moveList) != 0 {
				moves[i] = moveList
			}
		}
	}
	return moves
}

func (b *Board) GetLegalMoves(pieceName Piece, pos int) []int {
	moves := b.GetPartialMoves(pieceName, pos, false)

	var legalMoves []int
	for _, targetPos := range moves {
		tempBoard := b.Copy()
		targetPiece := tempBoard.WhichPieceExists(byte(targetPos))
		tempBoard.setPiece(getPiece(pieceName, tempBoard.Turn), targetPiece, byte(pos), byte(targetPos))
		if !tempBoard.IsKingInCheck(b.Turn) {
			legalMoves = append(legalMoves, targetPos)
		}
	}
	return legalMoves
}

func (b *Board) GetPartialMoves(pieceName Piece, pos int, attackMoves bool) []int {
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

func (b *Board) UpdateBoard(from byte, to byte) {
	pieceName, color := getNameAndColor(b.WhichPieceExists(from))
	targetPieceName, targetColor := getNameAndColor(b.WhichPieceExists(to))

	b.EnPassantSquare = 255
	b.HalfMove++
	if targetPieceName != 0 {
		b.HalfMove = 0
		b.Score -= targetPieceName.Value() * targetColor.Value()
	}

	switch pieceName {
	case Pawn:
		b.HalfMove = 0
		// Update en-passant squarePos
		if subtract(from, to) == 16 {
			b.EnPassantSquare = to
		}
		// Make en-passant Move
		if b.EnPassantSquare != 255 && subtract(to, b.EnPassantSquare) == 8 {
			b.Score -= targetColor.Value()
		}
		if b.isPawnPromoted(to) {
			promotionPiece := Queen
			b.Score += promotionPiece.Value() * b.Turn.Value()
		}
	case Rook:
		// setting casting bit to zero : binary(3) -> 11 && binary(12) -> 1100
		if b.Turn == White {
			b.Castle = b.Castle & ^byte(3)
		} else {
			b.Castle = b.Castle & ^byte(12)
		}
	case King:
		if b.Turn == White {
			b.Castle = b.Castle & ^byte(3)
		} else {
			b.Castle = b.Castle & ^byte(12)
		}
		// If castling move than also move the rook
		if subtract(from, to) > 1 {
			var currRook byte
			var targetRook byte
			if from > to {
				currRook = to - 2
				targetRook = to + 1
			} else {
				currRook = to + 1
				targetRook = to - 1
			}
			b.setPiece(getPiece(Rook, b.Turn), getPiece(0, 0), currRook, targetRook)
		}
	default:
	}
	b.setPiece(getPiece(pieceName, color), getPiece(targetPieceName, targetColor), from, to)
	b.rotateTurn()
}

func (b *Board) IsKingInCheck(color Color) bool {
	KingPos := b.getKingPosition(color)
	// todo change getallmoves to other function
	moveList := b.GetAllMoves(true)

	for _, moves := range moveList {
		if slices.Contains(moves, KingPos) {
			return true
		}
	}
	return false
}

func (b *Board) WhichPieceExists(pos byte) byte {
	// todo there is coming negative pos test it and solve it
	for piece, positions := range b.Bitboard {
		if positions&(1<<pos) != 0 {
			return piece
		}
	}
	return 0
}

func (b *Board) IsPieceExists(pieceName Piece, color Color, pos int) bool {
	return b.Bitboard[getPiece(pieceName, color)]>>pos&1 != 0
}

func (b *Board) DrawBoard(moves []int) {
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			pos := rank*8 + file
			piece := b.WhichPieceExists(byte(pos))
			if slices.Contains(moves, pos) {
				fmt.Printf("*\t")
				continue
			}
			fmt.Printf("%d\t", piece)

		}
		fmt.Println()
	}
}

func (b *Board) Copy() *Board {
	newBoard := &Board{
		Bitboard:        b.Bitboard,
		EnPassantSquare: b.EnPassantSquare,
		Turn:            b.Turn,
		HalfMove:        b.HalfMove,
		FullMove:        b.FullMove,
		Castle:          b.Castle,
		Score:           b.Score,
	}
	newBoard.Bitboard = make(map[byte]uint64)

	for piece, positions := range b.Bitboard {
		newBoard.Bitboard[piece] = positions
	}

	return newBoard
}

func (b *Board) knightMoves(pos int) []int {
	var moves []int
	offset := []int{-17, -15, -10, -6, 6, 10, 15, 17}

	currRank := getRank(pos)
	currFile := getFile(pos)
	for i := 0; i < len(offset); i++ {
		newPos := pos + offset[i]
		if !isOnBoard(newPos) {
			continue
		}

		newFile := getFile(newPos)
		newRank := getRank(newPos)

		if abs(currRank-newRank) > 2 || abs(currFile-newFile) > 2 {
			continue
		}
		_, color := getNameAndColor(b.WhichPieceExists(byte(newPos)))
		if color != b.Turn {
			moves = append(moves, newPos)
		}
	}
	return moves
}

func (b *Board) bishopMoves(pos int) []int {
	var moves []int
	offset := []int{-9, -7, 7, 9}

	for i := 0; i < len(offset); i++ {
		newPos := pos + offset[i]
		oldPos := pos
		for isOnBoard(newPos) && getRank(newPos) != getRank(oldPos) {
			_, color := getNameAndColor(b.WhichPieceExists(byte(newPos)))
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
	offset := []int{-1, 1, -8, 8}

	for i := 0; i < len(offset); i++ {
		newPos := pos + offset[i]
		oldPos := pos
		for isOnBoard(newPos) && (getRank(oldPos) == getRank(newPos) ||
			getFile(oldPos) == getFile(newPos)) {
			_, color := getNameAndColor(b.WhichPieceExists(byte(newPos)))

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
	offset := []int{7, 9}

	for i := 0; i < len(offset); i++ {
		newPos := pos + offset[i]*b.Turn.Value()
		if !isOnBoard(newPos) {
			continue
		}
		_, color := getNameAndColor(b.WhichPieceExists(byte(newPos)))
		if color != b.Turn && color != 0 {
			moves = append(moves, newPos)
		}
	}
	return moves
}

func (b *Board) pawnMoves(pos int) []int {
	var moves []int
	offset := 8 * b.Turn.Value()
	newPos := pos + offset
	if isOnBoard(newPos) {
		pieceName, _ := getNameAndColor(b.WhichPieceExists(byte(newPos)))
		if pieceName == 0 {
			moves = append(moves, newPos)
			color := b.Turn
			currentRank := getRank(pos)

			if (color == White && currentRank == 1) || (color == Black && currentRank == 6) {
				newPos += offset
				pieceName, _ := getNameAndColor(b.WhichPieceExists(byte(newPos)))
				if pieceName == 0 {
					moves = append(moves, newPos)
				}
			}
		}
	}

	// en-passant move
	if b.EnPassantSquare != 255 {
		if getRank(pos) == getRank(int(b.EnPassantSquare)) &&
			abs(getFile(pos)-getFile(int(b.EnPassantSquare))) == 1 {
			moves = append(moves, int(b.EnPassantSquare)+offset)
		}
	}
	return moves
}

func (b *Board) kingMoves(pos int) []int {
	var moves []int
	offset := []int{-9, -8, -7, -1, 1, 7, 8, 9}

	for i := 0; i < len(offset); i++ {
		newPos := pos + offset[i]
		if !isOnBoard(newPos) {
			continue
		}
		_, color := getNameAndColor(b.WhichPieceExists(byte(newPos)))
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
	isPossible := b.Castle >> bitshift & 1
	// todo check it later
	if isPossible != 0 {
		rookPos := pos - 4
		pathClear := true

		for newPos := rookPos + 1; newPos < pos; newPos++ {
			pieceName, _ := getNameAndColor(b.WhichPieceExists(byte(newPos)))
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
	// todo check it later
	isPossible := b.Castle >> bitshift & 1
	if isPossible != 0 {
		rookPos := pos + 3
		pathClear := true

		for newPos := pos + 1; newPos < rookPos; newPos++ {
			pieceName, _ := getNameAndColor(b.WhichPieceExists(byte(newPos)))
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

func (b *Board) isPawnPromoted(pos byte) bool {
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

func (b *Board) setPiece(piece, targetPiece byte, from, to byte) {
	// todo add nor operation to solve all cases
	if targetPiece != 0 {
		b.Bitboard[targetPiece] &= ^(1 << to)
	}

	b.Bitboard[piece] ^= 1 << from
	b.Bitboard[piece] ^= 1 << to
}

func (b *Board) getKingPosition(color Color) int {
	kingPos := b.Bitboard[getPiece(King, color)]
	pos := 0
	for ; kingPos != 1; pos++ {
		kingPos = kingPos >> 1
	}
	return pos
}

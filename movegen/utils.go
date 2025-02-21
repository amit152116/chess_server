package movegen

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Piece byte

const (
	Pawn Piece = iota + 1
	Knight
	Bishop
	Rook
	Queen
	King
)

func (p Piece) Notation() string {
	var notation string
	switch p {
	case Pawn:
		notation = "p"
	case Knight:
		notation = "n"
	case Bishop:
		notation = "b"
	case Rook:
		notation = "r"
	case Queen:
		notation = "q"
	case King:
		notation = "k"
	}
	return notation
}

func fromNotation(notation string) Piece {
	var pieceName Piece
	switch notation {
	case "p":
		pieceName = Pawn
	case "n":
		pieceName = Knight
	case "b":
		pieceName = Bishop
	case "r":
		pieceName = Rook
	case "q":
		pieceName = Queen
	case "k":
		pieceName = King

	}
	return pieceName
}

func (p Piece) Value() int {
	var value int
	switch p {
	case Pawn:
		value = 1
	case Knight:
		value = 3
	case Bishop:
		value = 3
	case Rook:
		value = 5
	case Queen:
		value = 9
	case King:
		value = 255
	}
	return value
}

func (p Piece) Name() string {
	var name string
	switch p {
	case Pawn:
		name = "Pawn"
	case Knight:
		name = "Knight"
	case Bishop:
		name = "Bishop"
	case Rook:
		name = "Rook"
	case Queen:
		name = "Queen"
	case King:
		name = "King"
	}
	return name
}

type Color byte

const (
	White Color = 8
	Black Color = 16
)

func (c Color) Value() int {
	var val int
	switch c {
	case White:
		val = 1
	case Black:
		val = -1
	}
	return val
}

func (c Color) Name() string {
	var name string
	switch c {
	case White:
		name = "White"
	case Black:
		name = "Black"
	}
	return name
}

func getPiece(p Piece, c Color) byte {
	return byte(c) | byte(p)
}

func getNameAndColor(piece byte) (Piece, Color) {
	pieceName := piece & 7
	return Piece(pieceName), Color(piece - pieceName)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func subtract(a, b byte) byte {
	if a > b {
		return a - b
	}
	return b - a
}

func getFile(pos int) int {
	return pos & 7
}

func getRank(pos int) int {
	return pos >> 3
}

func getSquareName(pos byte) string {
	if pos >= 64 || pos < 0 {
		return "-"
	}
	file := getFile(int(pos))
	rank := getRank(int(pos))
	return fmt.Sprintf("%c%d", rune(file+97), rank+1)
}

func isOnBoard(targetPos int) bool {
	if targetPos >= 64 || targetPos < 0 {
		return false
	}
	return true
}

func ValidateFEN(fen interface{}) bool {
	switch fen.(type) {
	case string:
		rules := strings.Split(fen.(string), " ")
		return ValidateFEN(rules)
	case []string:
		rules := fen.([]string)
		if len(rules) != 6 {
			return false
		}

		return validateBoard(rules[0]) &&
			validateTurnToMove(rules[1]) &&
			validateCastlingRule(rules[2]) &&
			validateEnpassant(rules[3]) &&
			validateHalfMove(rules[4]) &&
			validateFullMove(rules[5])
	default:
		return false
	}
}

func validateBoard(board string) bool {
	rows := strings.Split(board, "/")
	if len(rows) != 8 {
		return false
	}

	totalSquares := 0
	for _, row := range rows {
		squares := 0
		for _, char := range row {
			if char >= '1' && char <= '8' {
				squares += int(char - '0')
			} else if strings.ContainsRune("prnbqkPRNBQK", char) {
				squares++
			} else {
				return false
			}
		}
		if squares != 8 {
			return false
		}
		totalSquares += squares
	}

	if totalSquares != 64 {
		return false
	}

	return true
}

func validateTurnToMove(rule string) bool {
	regex := regexp.MustCompile(`^[wb]$`)
	if regex.MatchString(rule) {
		return true
	}
	return false
}

func validateEnpassant(rule string) bool {
	if rule == "-" {
		return true
	}
	regex := regexp.MustCompile(`^[a-h][36]$`)
	if regex.MatchString(rule) {
		return true
	}
	return false
}

func validateHalfMove(rule string) bool {
	moves, err := strconv.Atoi(rule)
	if err != nil {
		return false
	}
	if moves > 50 || moves < 0 {
		return false
	}
	return true
}

func validateFullMove(rule string) bool {
	moves, err := strconv.Atoi(rule)
	if err != nil {
		return false
	}
	if moves < 1 {
		return false
	}
	return true
}

func validateCastlingRule(rule string) bool {
	if rule == "" {
		return false
	}
	regex := regexp.MustCompile(`^(-|(K?Q?k?q?))$`)
	if regex.MatchString(rule) {
		return true
	}

	return false
}

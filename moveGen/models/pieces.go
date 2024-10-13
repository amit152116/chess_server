package models

type PieceName uint8

const (
	Pawn PieceName = iota + 1
	Knight
	Bishop
	Rook
	Queen
	King
)

func (p PieceName) Notation() string {
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

func fromNotation(notation string) PieceName {
	var pieceName PieceName
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
func (p PieceName) Value() int {
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

func (p PieceName) Name() string {
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

type Color uint8

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
func GetPiece(p PieceName, c Color) uint8 {
	return uint8(c) | uint8(p)
}

func GetNameAndColor(piece uint8) (PieceName, Color) {
	var pieceName = piece & 7
	return PieceName(pieceName), Color(piece - pieceName)
}

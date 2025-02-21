package utils

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const SessionTimeout = int64(60 * 30)

// GameStatus represents the status of a game
type GameStatus int

const (
	GameStatusInProgress GameStatus = iota
	GameStatusFinished
)

func (g GameStatus) String() string {
	return [...]string{"in_progress", "finished"}[g]
}

// TimeControl represents the time control of a game
type TimeControl int

const (
	Bullet TimeControl = iota + 1
	Blitz
	Rapid
	Classical
)

func (t TimeControl) String() string {
	return [...]string{"bullet", "blitz", "rapid", "classical"}[t]
}

// FriendStatus represents the status of a friend request
type FriendStatus int

const (
	Pending FriendStatus = iota
	Accepted
)

func (f FriendStatus) String() string {
	return [...]string{"pending", "accepted"}[f]
}

// GameResult represents the result of a game
type GameResult int

const (
	Checkmate GameResult = iota
	Draw
	Resignation
	Timeout
	Stalemate
)

func (g GameResult) String() string {
	return [...]string{"checkmate", "draw", "resignation", "timeout", "stalemate"}[g]
}

// SSLMode represents the SSL mode of a connection
type SSLMode int

const (
	SSLModeDisable SSLMode = iota
	SSLModeRequire
	SSLModeVerifyCA
	SSLModeVerifyFull
)

func (s SSLMode) String() string {
	return [...]string{"disable", "require", "verify-ca", "verify-full"}[s]
}

type Role int

const (
	RoleAdmin Role = iota
	RoleUser
	RoleGuest
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EmailValidation(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._+-]{5,30}@[a-zA-Z.]+\.[a-zA-Z]{2,}$`)

	return re.MatchString(email)
}

func LengthToBytes(length int) []byte {
	lengthBytes := []byte{byte(length >> 8), byte(length & 0xFF)}
	return lengthBytes
}

func BytesToLength(bytes []byte) uint16 {
	return uint16(bytes[0])<<8 | uint16(bytes[1])
}

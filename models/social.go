package models

import (
	"github.com/Amit152116Kumar/chess_server/utils"
	"time"
)

type ChatMessage struct {
	SenderID int       `json:"sender_id"`
	Message  string    `json:"message"`
	SentAt   time.Time `json:"sent_at"`
}

type Friends struct {
	UserID    int                `json:"user_id"`
	FriendID  int                `json:"friend_id"`
	CreatedAt time.Time          `json:"created_at"`
	Status    utils.FriendStatus `json:"status"`
}

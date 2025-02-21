package models

import (
	"github.com/Amit152116Kumar/chess_server/utils"
	"github.com/google/uuid"
	"time"
)

type Session struct {
	UID    uuid.UUID  `json:"uid" redis:"uid"`
	Email  string     `json:"email" redis:"email"`
	Expiry int64      `json:"expiry" redis:"expiry"`
	Role   utils.Role `json:"role" redis:"role"`
}

func NewSession(email string, role utils.Role) *Session {
	return &Session{
		UID:    uuid.New(),
		Email:  email,
		Expiry: time.Now().Unix() + utils.SessionTimeout,
		Role:   role,
	}
}

func (s *Session) IsValid() bool {
	return time.Now().Unix() < s.Expiry
}

func (s *Session) Refresh(uid uuid.UUID) {
	s.UID = uid
	s.Expiry = time.Now().Unix() + utils.SessionTimeout
}

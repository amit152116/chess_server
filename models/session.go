package models

import (
	"github.com/Amit152116Kumar/chess_server/utils"
	"github.com/google/uuid"
	"time"
)

type Session struct {
	UID    uuid.UUID
	email  string
	expiry int64
	Role   utils.Role
}

func NewSession(email string, role utils.Role) *Session {
	return &Session{
		UID:    uuid.New(),
		email:  email,
		expiry: time.Now().Unix() + utils.SessionTimeout,
		Role:   role,
	}
}

func (s *Session) IsValid() bool {
	return time.Now().Unix() < s.expiry
}

func (s *Session) Refresh(uid uuid.UUID) {
	s.UID = uid
	s.expiry = time.Now().Unix() + utils.SessionTimeout
}

var Sessions = make(map[uuid.UUID]*Session)

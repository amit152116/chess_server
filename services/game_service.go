package services

import (
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/google/uuid"
)

func CreateGame(timeControl *models.NewGameReqParam) uuid.UUID {
	uid, _ := uuid.NewV7()

	return uid
}

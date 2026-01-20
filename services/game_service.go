package services

import (
	"github.com/amit152116/chess_server/models"
	"github.com/google/uuid"
)

func CreateGame(timeControl *models.NewGameReqParam) uuid.UUID {
	uid, _ := uuid.NewV7()

	return uid
}

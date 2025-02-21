package models

import (
	"time"

	"github.com/Amit152116Kumar/chess_server/utils"
)

type Game struct {
	ID            int               `json:"id"`
	WhitePlayerID int               `json:"white_id"`
	BlackPlayerID int               `json:"black_id"`
	StartTime     time.Time         `json:"start_time"`
	EndTime       time.Time         `json:"end_time,omitempty"`
	Status        utils.GameStatus  `json:"status"`
	TimeControl   utils.TimeControl `json:"time_control"`
	WinnerID      int               `json:"winner_id,omitempty"`
}

type Move struct {
	GameID     int       `json:"game_id"`
	PlayerID   int       `json:"player_id"`
	MoveNumber int       `json:"move_number"`
	Move       string    `json:"move"`
	MoveTime   time.Time `json:"move_time"`
}

type NewGameReqParam struct {
	TimeControl utils.TimeControl `json:"time_control" form:"time_control" binding:"required"`
	Time        int               `json:"time" form:"time" binding:"required" `
	Increment   int               `json:"increment" form:"increment" binding:"required"`
	IsRandom    string            `json:"is_random" form:"is_random" `
}

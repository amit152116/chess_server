package models

import (
	"github.com/Amit152116Kumar/chess_server/utils"
	"time"
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

type NewGame struct {
	TimeControl utils.TimeControl `json:"time_control" form:"time_control" binding:"required"`
}

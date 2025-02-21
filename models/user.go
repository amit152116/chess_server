package models

import (
	"github.com/google/uuid"
	"time"
)

type RegisterUserPayload struct {
	Username  string `json:"username" binding:"required" form:"username"`
	Password  string `json:"password" binding:"required" form:"password"`
	Email     string `json:"Email" binding:"required" form:"Email"`
	FirstName string `json:"first_name" binding:"required" form:"first_name"`
	LastName  string `json:"last_name" binding:"required" form:"last_name"`
}

type LoginUserPayload struct {
	Email    string `json:"Email" binding:"required" form:"Email"`
	Password string `json:"password" binding:"required" form:"password"`
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"Email"`
	AvatarURL    string    `json:"avatar_url"`
	Bio          string    `json:"bio"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PasswordHash string    `json:"password"`
}

type Stats struct {
	UserDetails User `json:"user_details"`
	Bullet      int  `json:"bullet"`
	Blitz       int  `json:"blitz"`
	Rapid       int  `json:"rapid"`
	Classical   int  `json:"classical"`
	TotalGames  int  `json:"total_games"`
	Wins        int  `json:"wins"`
	Losses      int  `json:"losses"`
	Draws       int  `json:"draws"`
}

package services

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/db"
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/myErrors"
	"github.com/Amit152116Kumar/chess_server/redis"
	"github.com/Amit152116Kumar/chess_server/utils"
	"github.com/google/uuid"
)

var dbInstance = db.Instance

func RegisterUser(user *models.RegisterUserPayload) (uuid.UUID, error) {
	if !utils.EmailValidation(user.Email) {
		return uuid.Nil, myErrors.InvalidEmail
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		// todo: there may be a error in hashing password but Response is given as user already exists
		return uuid.Nil, err
	}
	user.Password = hash
	if err := dbInstance.AddUser(user); err != nil {
		return uuid.Nil, myErrors.UserAlreadyExists
	}
	session := models.NewSession(user.Email, utils.RoleUser)
	result, err := redis.Client.HSet(redis.Ctx, session.UID.String(), session).Result()
	if err != nil {
		return uuid.Nil, err
	}
	fmt.Println(result)
	return session.UID, nil
}

func AuthenticateUser(user *models.LoginUserPayload) (uuid.UUID, error) {

	hash, err := dbInstance.GetPassword(user.Email)
	if err != nil || !utils.CheckPasswordHash(user.Password, hash) {
		return uuid.Nil, myErrors.InvalidCredentials
	}

	session := models.NewSession(user.Email, utils.RoleUser)
	result, err := redis.Client.HSet(redis.Ctx, session.UID.String(), session).Result()
	if err != nil {
		return uuid.Nil, err
	}
	fmt.Println(result)
	return session.UID, nil
}

func RefreshToken(uid string) uuid.UUID {
	key := uuid.MustParse(uid)
	var session models.Session
	err := redis.Client.HGetAll(redis.Ctx, key.String()).Scan(&session)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil
	}
	newUid := uuid.New()
	session.Refresh(newUid)
	return newUid
}

func Guest() uuid.UUID {
	session := models.NewSession("guest", utils.RoleGuest)
	result, err := redis.Client.HSet(redis.Ctx, session.UID.String(), session).Result()
	if err != nil {
		return uuid.Nil
	}
	fmt.Println(result)
	return session.UID
}

package services

import (
	"github.com/Amit152116Kumar/chess_server/models"
)

func UpdateUser(user *models.User) error {
	if user.Email != "" {
		if err := dbInstance.UpdateEmail(user.Username, user.Email); err != nil {
			return err
		}
	}
	if user.Username != "" {

	}
	return nil
}

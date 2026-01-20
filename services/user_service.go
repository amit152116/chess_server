package services

import (
	"github.com/amit152116/chess_server/models"
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

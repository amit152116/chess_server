package test

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/services"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	t.Run("Test Register DBUser", func(t *testing.T) {
		user := models.RegisterUserPayload{
			Username:  "test",
			Password:  "test",
			Email:     "test",
			FirstName: "test",
			LastName:  "test",
		}
		_, err := services.RegisterUser(&user)
		if err != nil {
			t.Errorf("Error in Register DBUser")
		}

		uid, err := services.AuthenticateUser(&models.LoginUserPayload{
			Email: "test",
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(uid)

	})
}

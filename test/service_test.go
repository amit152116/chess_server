package test

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/services"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	t.Run("Test Register User", func(t *testing.T) {
		user := models.RegisterUserPayload{
			"test",
			"test",
			"test",
			"test",
			"test",
		}
		_, err := services.RegisterUser(&user)
		if err != nil {
			t.Errorf("Error in Register User")
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

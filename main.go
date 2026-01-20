package main

import (
	"log"

	"github.com/amit152116/chess_server/cmd"
	"github.com/amit152116/chess_server/config"
	"github.com/amit152116/chess_server/utils"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Panicln("connectToDB: ", err)
	}
	config.Cfg.SetSSLMode(utils.SSLModeDisable)
	cmd.Execute()
}

// todo Add more comments for better readability
//  Handle myErrors properly

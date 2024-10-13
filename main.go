package main

import (
	"github.com/Amit152116Kumar/chess_server/cmd"
	"github.com/Amit152116Kumar/chess_server/config"
	"github.com/Amit152116Kumar/chess_server/utils"
	"log"
)

func main() {
	err := config.LoadConfig("local")
	if err != nil {
		log.Panicln("connectToDB: ", err)
	}
	config.Cfg.SetSSLMode(utils.SSLModeDisable)
	cmd.Execute()
}

// todo Add more comments for better readability
//  Handle myErrors properly

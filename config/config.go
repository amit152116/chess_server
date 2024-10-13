package config

import (
	"encoding/json"
	"fmt"
	"github.com/Amit152116Kumar/chess_server/utils"
	"io"
	"os"
)

type DbConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

func (a *DbConfig) SetSSLMode(sslMode utils.SSLMode) {
	a.SSLMode = sslMode.String()
}
func (a *DbConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		a.Host, a.Port, a.User, a.Password, a.DBName, a.SSLMode)
}

var Cfg = &DbConfig{}

func LoadConfig(postgresProvider string) error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var jsonData []byte
	if jsonData, err = io.ReadAll(file); err != nil {
		return err
	}

	var data map[string]json.RawMessage
	if err = json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	postgres := DbConfig{}
	if err = json.Unmarshal(data[postgresProvider], &postgres); err != nil {
		return err
	}
	*Cfg = postgres
	return nil
}

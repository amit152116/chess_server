package config

import (
	"fmt"
	"os"

	"github.com/amit152116/chess_server/utils"
	"github.com/joho/godotenv"
)

type Configs struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	SSLMode       string
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func (a *Configs) SetSSLMode(sslMode utils.SSLMode) {
	a.SSLMode = sslMode.String()
}

func (a *Configs) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		a.DBHost, a.DBPort, a.DBUser, a.DBPassword, a.DBName, a.SSLMode)
}

var Cfg *Configs

func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	Cfg = &Configs{
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBName:        os.Getenv("DB_NAME"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
	return nil
}

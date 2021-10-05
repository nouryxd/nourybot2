package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Username string
	Oauth    string
	DB_Name  string
	DB_User  string
	DB_Pass  string
	DB_Host  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	cfg := &Config{
		Username: os.Getenv("TWITCH_USER"),
		Oauth:    os.Getenv("TWITCH_PASS"),
		DB_Name:  os.Getenv("DB_NAME"),
		DB_User:  os.Getenv("DB_USER"),
		DB_Pass:  os.Getenv("DB_PASS"),
		DB_Host:  os.Getenv("DB_HOST"),
	}
	log.Info("Config loaded succesfully")

	return cfg

}

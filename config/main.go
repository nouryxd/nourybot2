package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Username string
	Oauth    string
	Client   string
	Secret   string
	DB_Name  string
	DB_User  string
	DB_Pass  string
	DB_Host  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}
	fmt.Println("loadconfig")
	cfg := &Config{
		Username: os.Getenv("TWITCH_USER"),
		Oauth:    os.Getenv("TWITCH_PASS"),
		Client:   os.Getenv("TWITCH_CLIENT"),
		Secret:   os.Getenv("TWITCH_SECRET"),
		DB_Name:  os.Getenv("DB_NAME"),
		DB_User:  os.Getenv("DB_USER"),
		DB_Pass:  os.Getenv("DB_PASS"),
		DB_Host:  os.Getenv("DB_HOST"),
	}
	// fmt.Println(cfg)
	return cfg
}

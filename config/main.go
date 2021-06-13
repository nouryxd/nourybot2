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
	}
	fmt.Println(cfg)
	return cfg
}

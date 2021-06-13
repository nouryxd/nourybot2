package config

import (
	"fmt"
	"os"
)

type Config struct {
	Username string
	Oauth    string
}

// func NewConfig() *Config {
// 	fmt.Println("newconfig")
// 	cfg := LoadConfig()
// 	return cfg
// }

func LoadConfig() *Config {
	fmt.Println("loadconfig")
	cfg := &Config{
		Username: "nourybot",
		Oauth:    string(os.Getenv("TWITCH_NOURYBOT_OAUTH")),
	}
	return cfg
}

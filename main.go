package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/config"
)

func main() {
	cfg := config.LoadConfig()
	client := twitch.NewClient(cfg.Username, cfg.Oauth)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})
	client.Join("nouryqt")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
	client.Connect()
}

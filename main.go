package main

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/config"
)

func main() {
	cfg := config.LoadConfig()

	client := twitch.NewClient(cfg.Username, cfg.Oauth)

	// client.OnPrivateMessage(func(message twitch.PrivateMessage) {
	// 	handlers.HandlePrivateMessage(message)
	// })
	bot := bot.NewBot(cfg, client)
	bot.Connect("nouryqt")

}

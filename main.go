package main

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/config"
)

func main() {
	// Load config containing passwords
	cfg := config.LoadConfig()

	// Create a new twitch with parameters specified from
	// the config module
	client := twitch.NewClient(cfg.Username, cfg.Oauth)

	bot := bot.NewBot(cfg, client)
	bot.Connect()

}

package main

import (
	"log"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/config"
)

func main() {
	// Load config containing client configurations
	cfg := config.LoadConfig()

	// Create a new twitch with parameters specified from
	// the config module
	client := twitch.NewClient(cfg.Username, cfg.Oauth)

	// Creat New Bot with twitch client and
	// config and connect to chat
	bot := bot.NewBot(cfg, client)
	err := bot.Connect()
	if err != nil {
		log.Fatal("Couldn't establish connection", err)
	}

}

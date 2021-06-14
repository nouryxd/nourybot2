package main

import (
	"log"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/config"
	"github.com/lyx0/nourybot-go/db"
)

func main() {
	// Load config containing client configurations
	cfg := config.LoadConfig()

	// Create a new twitch with parameters specified from
	// the config module
	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)

	sqlClient := db.Connect()

	// Creat New Bot with twitch client and
	// config and connect to chat
	bot := bot.NewBot(cfg, twitchClient, sqlClient)
	err := bot.Connect()
	if err != nil {
		log.Fatal("Couldn't establish connection", err)
	}

}

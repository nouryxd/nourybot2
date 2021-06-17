package client

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot2/handlers"
	log "github.com/sirupsen/logrus"
)

func DiscordConnect() {
	DC_AUTH := os.Getenv("DC_AUTH")
	dcb, err := discordgo.New("Bot " + DC_AUTH)
	if err != nil {
		log.Fatal("Error authenticating with discord", err)
	}

	// Register funcs as callbacks for events.
	dcb.AddHandler(
		handlers.DiscordMessage,
	)

	dcb.Identify.Intents = discordgo.IntentsGuildMessages

	err = dcb.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	log.Info("Connected to Discord!")
}

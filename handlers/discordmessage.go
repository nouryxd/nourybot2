package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func DiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info(m.Message)
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

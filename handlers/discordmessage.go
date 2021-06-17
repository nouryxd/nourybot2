package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot2/commands"
)

func DiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Don't act on our own messages
	if m.Message.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!ping" {
		commands.Ping(s, m)
	}
	if m.Content == "!pong" {
		commands.Pong(s, m)
	}
}

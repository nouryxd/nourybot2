package bot

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	twitch "github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot2/config"
	"github.com/lyx0/nourybot2/handlers"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	cfg           *config.Config
	twitchClient  *twitch.Client
	discordClient *discordgo.Session
	sqlClient     *sql.DB
}

func NewBot(cfg *config.Config, twitchClient *twitch.Client, discordClient *discordgo.Session, sqlClient *sql.DB) *Bot {
	return &Bot{
		cfg:           cfg,
		twitchClient:  twitchClient,
		discordClient: discordClient,
		sqlClient:     sqlClient,
	}
}

func (b *Bot) newTwitchClient() *twitch.Client {
	twitchClient := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return twitchClient
}

func (b *Bot) newDiscordClient() *discordgo.Session {
	discordClient, err := discordgo.New("Bot " + b.cfg.DC_AUTH)
	if err != nil {
		log.Fatal("Error authenticating with discord", err)
	}

	return discordClient
}

func (b *Bot) discordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// if m.Author.ID == s.State.User.ID {
	// 	return
	// }
	handlers.DiscordMessage(s, m)
}

func (b *Bot) ConnectTwitch() error {
	log.Info("xd")
	twitchClient := b.newTwitchClient()

	twitchClient.Join("nouryqt")
	twitchClient.Say("nouryqt", "xd")

	twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		b.twitchMessage(message, twitchClient)
	})

	twitchClient.OnWhisperMessage(func(whisper twitch.WhisperMessage) {
		b.twitchWhisper(whisper, twitchClient)
	})

	err := twitchClient.Connect()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info("Connected to Twitch")
	return err
}

func (b *Bot) ConnectDiscord() error {
	log.Info("xd")

	discordClient := b.newDiscordClient()

	discordClient.AddHandler(b.discordMessage)

	discordClient.Identify.Intents = discordgo.IntentsGuildMessages

	err := discordClient.Open()
	if err != nil {
		log.Fatal("Error connecting to Discord: ", err)
	}

	log.Info("Connected to Discord.")
	return err
}

func (b *Bot) twitchMessage(message twitch.PrivateMessage, tc *twitch.Client) {
	handlers.TwitchMessage(message, tc)
}

// func (b *Bot) handleCommand(message twitch.PrivateMessage) {
// 	commands.HandleCommand(message)
// }

func (b *Bot) twitchWhisper(whisper twitch.WhisperMessage, tc *twitch.Client) {
	// handlers.twitchWhisper(whisper)
	handlers.TwitchWhisper(whisper, tc)
	log.Info(whisper)
}

func (b *Bot) closeConnection() {
	b.twitchClient.Disconnect()
	b.discordClient.Close()
}

// func (b *Bot) Say(channel string, message string) {
// 	b.twitchClient.Say(channel, message)
// }

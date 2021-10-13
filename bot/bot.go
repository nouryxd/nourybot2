package bot

import (
	twitch "github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot2/config"

	"github.com/lyx0/nourybot2/handlers"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	cfg          *config.Config
	twitchClient *twitch.Client
}

func NewBot(cfg *config.Config, twitchClient *twitch.Client) *Bot {
	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
	}
}

func (b *Bot) newTwitchClient() *twitch.Client {
	twitchClient := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return twitchClient
}

func (b *Bot) ConnectTwitch() error {
	log.Info("xd")
	twitchClient := b.newTwitchClient()

	twitchClient.Join("nourybot")
	twitchClient.Say("nourybot", "xd")

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

func (b *Bot) CloseConnection() {
	b.twitchClient.Disconnect()
}

// func (b *Bot) Say(channel string, message string) {
// 	b.twitchClient.Say(channel, message)
// }

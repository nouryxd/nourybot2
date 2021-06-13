package bot

import (
	"fmt"
	"time"

	twitch "github.com/gempir/go-twitch-irc/v2"
	config "github.com/lyx0/nourybot-go/config"
)

type Bot struct {
	cfg          *config.Config
	twitchClient *twitch.Client
	uptime       time.Time
}

func NewBot(cfg *config.Config, twitchClient *twitch.Client) *Bot {
	fmt.Println("newbot")
	return &Bot{
		cfg:          config.LoadConfig(),
		twitchClient: twitch.NewClient(cfg.Username, cfg.Oauth),
		uptime:       time.Now(),
	}
	// err := twitchClient.Connect()
	// if err != nil {
	// 	fmt.Printf("Error connecting to Twitch chat %#v", err)
	// }
}

func NewClient() twitch.Client {
	cfg := config.LoadConfig()
	client := twitch.NewClient(cfg.Username, cfg.Oauth)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})
	return *client
}

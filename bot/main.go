package bot

import (
	"fmt"
	"log"

	twitch "github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/config"
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
	// err := twitchClient.Connect()
	// if err != nil {
	// 	fmt.Printf("Error connecting to Twitch chat %#v", err)
	// }
}

func (b *Bot) newClient() *twitch.Client {
	fmt.Println("newClient")
	client := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	fmt.Println(b.cfg.Username, b.cfg.Oauth)
	client.OnPrivateMessage(b.handleMessage)
	client.Say("nouryqt", "HeyGuys")
	return client
}

func (b *Bot) Connect() {
	fmt.Println("Connect")
	client := b.newClient()
	b.Join()
	log.Fatal(client.Connect())

}

func (b *Bot) Join() {
	client := b.newClient()
	client.Join("nouryqt")
	fmt.Println("Joined nouryqt")
}

func (b *Bot) Say(channel string, text string) {
	b.twitchClient.Say(channel, text)
}

func (b *Bot) handleMessage(message twitch.PrivateMessage) {
	fmt.Println(message)
	fmt.Println("message xd")
}

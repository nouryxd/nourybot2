package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/lyx0/nourybot-go/bot"
)

var channels = map[string]*Channel{
	"nouryqt": {Name: "nouryqt"},
	"nrybot":  {Name: "nrybot"},
}

func connectToChannels() {
	for i := range channels {
		Nourybot.Client.Join(i)
		Nourybot.Client.Say(i, "FeelsDankMan")
	}
}

//func (c *Client) OnConnect(callback func()) {
//	SendTwitchMessage("nouryqt", "hehe")
//}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	botUser := os.Getenv("TWITCH_USER")
	botPass := os.Getenv("TWITCH_PASSWORD")
	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)

	Nourybot = &Bot{
		Client:   twitch.NewClient(botUser, botPass),
		Channels: channels,
	}

	err := Nourybot.Client.Connect()

	if err != nil {
		log.Fatalln(err.Error())
	}
	i

}

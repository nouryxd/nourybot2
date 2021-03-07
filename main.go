package main

import (
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	. "github.com/lyx0/nourybot-go/bot"
)

var channels = map[string]*Channel{
	"nouryqt": {Name: "nouryqt"},
	"nrybot":  {Name: "nrybot"},
}

func connectToChannels() {
	for i := range channels {
		Nourybot.Client.Join(i)
		Nourybot.Client.Say(i, "FeelsDankMan")
		log.Printf("Connected to channel: %v\n", i)
	}
}

func main() {
	log.Println("Starting")
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	botUser := os.Getenv("TWITCH_USER")
	botPass := os.Getenv("TWITCH_PASSWORD")

	Nourybot = &Bot{
		Client:   twitch.NewClient(botUser, botPass),
		Channels: channels,
	}

	// connectToChannels needs to be above err := Nourybot.Client.Connect()
	connectToChannels()

	err := Nourybot.Client.Connect()

	if err != nil {
		log.Fatalln(err.Error())
	}

}

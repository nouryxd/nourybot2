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

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	botUser := os.Getenv("TWITCH_USER")
	botPass := os.Getenv("TWITCH_PASSWORD")
	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	client := twitch.NewClient(botUser, botPass)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		if message.Message == "FeelsDankMan TeaTime" {
			client.Say("nouryqt", "pajaDink")
		}
	})

	client.Say("nouryqt", "pajaDink")

	client.Join("nouryqt")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

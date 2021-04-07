package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	. "github.com/lyx0/nourybot-go/bot"
	. "github.com/lyx0/nourybot-go/commands"
	db "github.com/lyx0/nourybot-go/mongo"
)

var channels = map[string]*Channel{
	"nouryqt":              {Name: "nouryqt"},
	"nourybot":             {Name: "nrybot"},
	"uudelleenkytkeytynyt": {Name: "uudelleenkytkeytynyt"},
	"xnoury":               {Name: "xnoury"},
	"nrybot":               {Name: "nrybot"},
	"noemience":            {Name: "noemience"},
}

func connectToChannels() {
	for i := range channels {
		Nourybot.Client.Join(i)
		SendTwitchMessage("nouryqt", "pajaDink")
		SendTwitchMessage("nourybot", ":)")
		SendTwitchMessage("nrybot", ":)")
		SendTwitchMessage("uudelleenkytkeytynyt", ":)")
		SendTwitchMessage("xnoury", "pajaDink")
		log.Printf("Connected to channel: %v\n", i)
	}
}

func main() {
	log.Println("Starting")

	mongoClient := db.Connect()

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	botUser := os.Getenv("TWITCH_USER")
	botPass := os.Getenv("TWITCH_PASSWORD")

	Nourybot = &Bot{
		Client:   twitch.NewClient(botUser, botPass),
		Mongo:    mongoClient,
		Channels: channels,
		// Commands: commands.initCommands(),
	}

	Nourybot.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		channelID := message.Tags["room-id"]
		if channelID == "" {
			fmt.Printf("Missing room-id tag in message")
			return
		}
		if message.Tags["user-id"] == "596581605" {
			return
		}
		fmt.Printf("%v\n", message.Message)
		HandleMessage(message, Nourybot)

	})

	// connectToChannels needs to be above err := Nourybot.Client.Connect()
	connectToChannels()
	// log.Print(mongoClient)
	err := Nourybot.Client.Connect()

	if err != nil {
		log.Fatalln(err.Error())
	}

}

package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

// https://random.dog/woof.json
type RandomDogResponse struct {
	Url string `json:"url"`
}

func HandleRandomDog(channel string) {
	response, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseObject RandomDogResponse
	json.Unmarshal(responseData, &responseObject)

	bot.SendTwitchMessage(channel, string(responseObject.Url))
}

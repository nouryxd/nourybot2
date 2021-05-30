package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

func HandleSubage(channel string, username string, streamer string) {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/twitch/subage/%s/%s", username, streamer))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject SubageResponse
	json.Unmarshal(body, &responseObject)

	if responseObject.SubageHidden {
		bot.SendTwitchMessage(channel, fmt.Sprintf("User "+username+" has their subscription status hidden. FeelsBadMan"))
	} else {
		months := fmt.Sprint(responseObject.Cumulative.Months)
		bot.SendTwitchMessage(channel, fmt.Sprintf("User "+username+" has been subscribed to "+streamer+" for "+months+" months."))
	}
}

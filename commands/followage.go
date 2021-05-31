package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

// https://api.ivr.fi
type followageApiResponse struct {
	User       string `json:"user"`
	UserID     string `json:"userid"`
	Channel    string `json:"channel"`
	ChannelId  string `json:"channelid"`
	FollowedAt string `json:"followedAt"`
	Error      string `json:"error"`
}

func HandleFollowage(channel string, username string, streamer string) {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/twitch/subage/%s/%s", username, streamer))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject followageApiResponse
	json.Unmarshal(body, &responseObject)

	// User or channel was not found
	if responseObject.Error != "" {
		bot.SendTwitchMessage(channel, fmt.Sprintf(responseObject.Error+" FeelsBadMan"))
		return
	} else if responseObject.FollowedAt == "" {
		bot.SendTwitchMessage(channel, fmt.Sprintf(username+" is not following "+streamer))
	} else {
		// Get followdate and trim the hours/minutes/seconds afterwards
		// TODO: Make it a nicer format, right now it's YYYY-MM-DD
		d := responseObject.FollowedAt[:10]
		bot.SendTwitchMessage(channel, fmt.Sprintf(username+" has been following "+streamer+" since "+d+"."))
	}
}

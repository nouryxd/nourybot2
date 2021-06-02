package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

func StrGenerateRandomNumber(max string) int {
	num, err := strconv.Atoi(max)
	if num < 1 {
		return 0
	}

	if err != nil {
		fmt.Printf("Supplied value %v is not a number", num)
		return 0
	} else {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(num)
	}
}

func GenerateRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// ElevatedPrivsMessage is a message from either moderator, vip or the broadcaster.
func ElevatedPrivsMessage(message twitch.PrivateMessage) bool {
	if message.User.Badges["moderator"] == 1 ||
		message.User.Badges["vip"] == 1 ||
		message.User.Badges["broadcaster"] == 1 {
		return true
	} else {
		return false
	}
}

// ModPrivsMessage is a message from either a moderator or the broadcaster but not vip.
func ModPrivsMessage(message twitch.PrivateMessage) bool {
	if message.User.Badges["moderator"] == 1 ||
		message.User.Badges["broadcaster"] == 1 {
		return true
	} else {
		return false
	}
}

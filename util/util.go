package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

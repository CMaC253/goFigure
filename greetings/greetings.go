package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//returns greeting for named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	//return greeting that embeds name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, v%. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]

}

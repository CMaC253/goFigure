package greetings

import "fmt"

//returns greeting for named person
func Hello(name string) string {
	//return greeting that embeds name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

package main

import (
	"fmt"
	"log"

	"github.com/cmac253/gofigure/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	//Get a greeting message and print it
	message, err := greetings.Hello("Dookie")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

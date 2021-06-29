package main

import (
	"fmt"
	"log"

	"github.com/cmac253/gofigure/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Sam", "Dobber"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

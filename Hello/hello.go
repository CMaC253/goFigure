package main

import (
	"fmt"

	"github.com/cmac253/gofigure/greetings"
)

func main() {
	//Get a greeting message and print it
	message := greetings.Hello("Meryl")
	fmt.Println(message)
}

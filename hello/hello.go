package main

import (
	"fmt"

	"github.com/MarQG/go-modules/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Ferenc")
	fmt.Println(message)
}
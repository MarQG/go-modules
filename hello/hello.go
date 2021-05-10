package main

import (
	"fmt"
	"log"

	"github.com/MarQG/go-modules/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	names := []string{"Ferenc", "Link", "Neo"}

	// If an error was returned, print it to the console and
	// exit the program.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
}
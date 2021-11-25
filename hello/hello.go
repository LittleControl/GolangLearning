package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix(("greetings: "))
	log.SetFlags(0)
	names := []string{"Gladys", "Mamnatha", "Darrin"}
	// message, err := greetings.Hello("Nothing")
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

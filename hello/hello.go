package main

import (
	"fmt"

	"example.com/greetings"

	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Artur")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Name1", "Name2", "Name3"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, name := range messages {
		fmt.Println(name)
	}
}

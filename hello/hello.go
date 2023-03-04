package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Artur")
	fmt.Println(message)
}

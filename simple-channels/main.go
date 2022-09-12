package main

import (
	"fmt"
	"strings"
)

func shout(ping, pong chan string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter q to quit)")

	for {
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}
		ping <- userInput
		// wait for response
		res := <-pong
		fmt.Println("Response:", res)
	}
	close(ping)
	close(pong)

	fmt.Println("All done. Closing channels.")
}

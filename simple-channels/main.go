package main

import (
	"fmt"
	"strings"
)

// listens to the channels
// ping is a receive-only channel, and pong is a send-only channel as indicated below
// these should be looked at from the POV of main when declared, i.e. in main ping always receives, pong always sends,
// so therefore: ping -> receive-only, pong -> send-only
func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping // assign s as everything received from ping
		if !ok {
			// do something
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s)) // send to pong formatted s
	}
}

func main() {
	// create two channels that only accept strings
	ping := make(chan string)
	pong := make(chan string)

	// start the shout function as a separate goroutine in the background and run forever
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		// send userInput to ping channel
		ping <- userInput
		// wait for a response on the pong channel
		response := <-pong // response is whatever comes in from pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}

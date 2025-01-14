package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10                  // how many clients the barbershop can seat at a time
var arrivalRate = 100                     // how often clients arrive to the barbershop (milliseconds)
var cutDuration = 1000 * time.Millisecond // how long it takes for the barber to cut one client
var timeOpen = 10 * time.Second           // how long the barbershop stays open

func main() {
	// print welcome message
	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity) // chan to store clients
	doneChan := make(chan bool)                      // chan to send a bool value when everything is done and to close down

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Obama")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen) // when you receive a signal from the channel that gets created from time.After(timeOpen), do:
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1
	go func() {
		for {
			// get a random number with average arrival rate
			randomMillseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block (keep the application going until things are finished) until you receive a closed signal to main
	<-closed
}

package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(time.Second * 6)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(time.Second * 3)
		ch <- "This is from server 2"
	}
}

func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("Case one:", s1)
		case s2 := <-channel1:
			fmt.Println("Case two:", s2)
		case s3 := <-channel2:
			fmt.Println("Case three:", s3)
		case s4 := <-channel2:
			fmt.Println("Case four:", s4)
			// default:
			// 	// avoiding deadlock
		}
		// If there's more than one case that matches functionally, the case will be selected randomly. I.e. s1 & s2, s3 & s4
	}
}

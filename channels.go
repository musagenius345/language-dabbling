package main

import (
	"fmt"
	"time"
)

func writeToChannel(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to the channel
		time.Sleep(time.Millisecond * 500)
	}
	close(ch) // Close the channel after sending data
}

func main() {
	ch := make(chan int) // Create an integer channel

	go writeToChannel(ch)

	// Receive data from the channel
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

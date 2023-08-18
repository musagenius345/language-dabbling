package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		fmt.Println("Letter:", string(char))
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	// Keep the main function running
	var input string
	fmt.Scanln(&input)
}

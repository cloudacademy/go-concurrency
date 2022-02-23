package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\nChannel Timeout Example...\n")

	channel := make(chan string)

	go func(channel chan string) {
		time.Sleep(5 * time.Second)
		channel <- "cloudacademy"

	}(channel)

	select {
	case msg1 := <-channel:
		fmt.Println(msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}

package main

import (
	"fmt"
	"time"
)

func in(channel chan<- string, msg string) {
	channel <- msg
}

func out(channel <-chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func main() {
	fmt.Printf("\nChannel Basic Example...\n")

	msgChan := make(chan string)

	go func() {
		msgChan <- "Cloud"
		msgChan <- "Academy"
		msgChan <- "2020"
	}()

	msg1 := <-msgChan
	msg2 := <-msgChan
	msg3 := <-msgChan

	fmt.Println(msg1, msg2, msg3)

	channel := make(chan string, 1)

	go out(channel)

	for i := 0; i < 10; i++ {
		in(channel, fmt.Sprintf("cloudacademy - %d", i))
	}

	time.Sleep(time.Second * 2) //crude
}

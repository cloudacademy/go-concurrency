package main

import (
	"bytes"
	"fmt"
)

func process(work <-chan string, fin chan<- string) {
	var b bytes.Buffer
	for {
		if msg, notClosed := <-work; notClosed {
			fmt.Printf("%s received...\n", msg)
			b.WriteString(msg)
		} else {
			fmt.Println("channel closed")
			fin <- b.String()
			return
		}
	}
}

func main() {
	fmt.Printf("\nChannel Close Example...\n")

	work := make(chan string, 3)
	fin := make(chan string)

	go process(work, fin)

	word := "cloudacademy"

	for j := 0; j < len(word); j++ {
		letter := string(word[j])
		work <- letter
		fmt.Printf("%s sent...\n", letter)
	}

	close(work)

	fmt.Printf("result: %s", <-fin)
}

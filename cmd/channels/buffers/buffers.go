package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\nChannel Buffers Example...\n")

	size := 3
	var buffChan = make(chan int, size)

	//reader
	go func() {
		for {
			<-buffChan
			time.Sleep(1 * time.Second)
		}
	}()

	//writer
	writer := func() {
		for i := 1; i <= 10; i++ {
			buffChan <- i
			println(i)
		}
	}

	writer()
}

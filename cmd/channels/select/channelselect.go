package main

import "fmt"

func main() {
	odds := make(chan int)
	evens := make(chan int)
	done := make(chan int)

	go func(ch, done chan int, data []int) {
		for _, v := range data {
			ch <- v
		}
		done <- 1
	}(odds, done, []int{1, 3, 5, 7, 9})

	go func(ch, done chan int, data []int) {
		for _, v := range data {
			ch <- v
		}
		done <- 1
	}(evens, done, []int{2, 4, 6, 8, 10})

	for n := 2; n > 0; {
		select {
		case num := <-odds:
			fmt.Printf("odd number: %d\n", num)
		case num := <-evens:
			fmt.Printf("even number: %d\n", num)
		case <-done:
			n--
		}
	}
}

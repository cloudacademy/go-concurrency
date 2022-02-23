package main

import "fmt"

func main() {
	odds := make(chan int)
	evens := make(chan int)
	done := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func(ch chan int, done chan int, data []int) {
		for _, n := range data {
			if n%2 != 0 {
				ch <- n
			}
		}
		done <- 1
	}(odds, done, numbers)

	go func(ch chan int, done chan int, data []int) {
		for _, n := range data {
			if n%2 == 0 {
				ch <- n
			}
		}
		done <- 1
	}(evens, done, numbers)

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

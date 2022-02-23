package main

import (
	"fmt"
	"time"
)

func countdown(activity int) {
	for i := activity; i >= 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("blastoff")
}

func main() {
	fmt.Printf("\nGoroutines Example...\n")

	name := "cloudacademy"

	go countdown(10)

	sayBlah := func() {
		fmt.Println("blah")
	}

	for i := 1; i <= 5; i++ {
		go sayBlah()
	}

	go func() {
		fn := "f1"
		fmt.Printf("%s, anonymous function called...\n", fn)
	}()

	go func(data string) {
		fn := "f2"
		fmt.Printf("%s, anonymous function called...\n", fn)
		fmt.Printf("%s, %s\n", fn, data)
		fmt.Printf("%s, %s\n", fn, name)

	}("some data")

	time.Sleep(100 * time.Millisecond)
	fmt.Println("finished!")
}

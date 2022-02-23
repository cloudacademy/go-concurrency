package main

import (
	"fmt"
	"sync"
)

func doSomething(waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	fmt.Println("doing something")
}

func main() {
	fmt.Printf("\nWaitGroups Example...\n")

	fmt.Println("starting...")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go doSomething(&waitgroup)

	waitgroup.Add(1)
	go func() {
		defer waitgroup.Done()
		fmt.Println("anonymous goroutine")
	}()

	waitgroup.Wait()

	fmt.Println("finished")
}

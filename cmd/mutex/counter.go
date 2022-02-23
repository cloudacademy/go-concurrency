package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var y = 0

func incrementSafe(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x++
	m.Unlock()
	wg.Done()
}

func incrementUnsafe(wg *sync.WaitGroup) {
	y++
	wg.Done()
}

func countSafe() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 10000000; i++ {
		w.Add(1)
		go incrementSafe(&w, &m)
	}

	w.Wait()
}

func countUnsafe() {
	var w sync.WaitGroup

	for i := 0; i < 10000000; i++ {
		w.Add(1)
		go incrementUnsafe(&w)
	}

	w.Wait()
}

func main() {
	fmt.Printf("\nMutex Example...\n")

	start := time.Now()

	countSafe()
	countUnsafe()

	fmt.Printf("x count: %d\n", x)
	fmt.Printf("y count: %d\n", y)

	fmt.Printf("execution time %vms\n", time.Since(start).Seconds())
}

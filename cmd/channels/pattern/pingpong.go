package main

import (
	"fmt"
	"sync"
	"time"
)

type ball struct {
	hit int
}

func player(name string, table chan *ball, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ball := <-table
		ball.hit++
		fmt.Println(name, ball.hit)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	fmt.Printf("\nChannel PingPong Example...\n")

	var wg sync.WaitGroup
	wg.Add(2)

	table := make(chan *ball, 1)

	go player("playerA", table, &wg)
	go player("playerB", table, &wg)

	table <- &ball{}
	wg.Wait()
}

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://cloudacademy.com",
	"https://go.dev",
}

func fetcher(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s - %s\n", resp.Status, url)
}

func main() {
	fmt.Printf("\nWaitGroups HTTP Example...\n")

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetcher(url, &wg)
	}

	wg.Wait()
	fmt.Println()
}

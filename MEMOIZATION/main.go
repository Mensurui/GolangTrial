package main

import (
	"example.com/memoization/memo1"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	m := memo1.New(httpGetBody)
	//	for url := range incomingURLs() {
	//		start := time.Now()
	//		value, err := m.Get(url)
	//		if err != nil {
	//			log.Print(err)
	//			continue
	//		}
	//		fmt.Printf("%s, %s, %d bytes\n",
	//			url, time.Since(start), len(value.([]byte)))
	//	}

	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	urls := make(chan string)

	// Launch a goroutine to send URLs into the channel.
	go func() {
		defer close(urls)

		// List of example URLs; you can replace this with dynamic data.
		urlList := []string{
			"https://www.example.com",
			"https://www.google.com",
			"https://www.golang.org",
			"https://www.github.com",
		}

		// Send each URL into the channel.
		for _, url := range urlList {
			urls <- url
			time.Sleep(1 * time.Second) // Simulate delay between URL fetching
		}
	}()

	return urls
}

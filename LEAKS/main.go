package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		defer fmt.Println("doWork exited")
		defer close(terminated)
		go func() {
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	name := "Mensur"
	nameChan := make(chan string)

	terminated := doWork(done, nameChan)

	go func() {
		nameChan <- name // Send name to the channel
		fmt.Println("Sent name to nameChan")
	}()

	go func() {
		time.Sleep(1 * time.Minute) // Allow doWork to run for a bit
		fmt.Println("Canceling doWork goroutine")
		close(done) // Signal done
	}()

	<-terminated
	fmt.Println("Done")
}

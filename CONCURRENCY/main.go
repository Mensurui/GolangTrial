package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine one")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine two")
	}()

	hello := func(wg *sync.WaitGroup, i int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", i)
	}
	const numGreeters = 5

	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
	fmt.Println("All done")
}

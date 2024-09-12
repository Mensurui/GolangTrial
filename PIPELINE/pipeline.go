package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ { // Stop after 10 iterations
			naturals <- x
		}
		close(naturals) // Close naturals when done
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares) // Close squares when done
	}()

	// Printer (in main goroutine)
	for x := range squares { // Exit loop when squares is closed
		fmt.Println(x)
	}
}

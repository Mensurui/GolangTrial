package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
	}()

	stream, ok := <-stringStream
	if !ok {
		fmt.Println("Deadlock")
	}

	fmt.Println(stream)
}

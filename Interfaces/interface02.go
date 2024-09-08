package main

import (
	"example.com/interfaces/interfaces"
	"fmt"
)

func main() {
	var c interfaces.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dorthy"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	interfaces.Sleep()
}

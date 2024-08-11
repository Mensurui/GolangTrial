package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, args := range os.Args[1:] {
		s += args + sep
		sep = " "
		fmt.Println(index)
	}
	fmt.Println(s)

}

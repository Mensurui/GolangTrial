package main

import "fmt"

func main() {
	fmt.Println("vim-go\n")
	fmt.Println("vim-don'tgo? \"")
	const GoUsage = `Go is a tool used for managing GO source code

	Usage:
		go command [arguments]
	...`
	fmt.Println(GoUsage)
}

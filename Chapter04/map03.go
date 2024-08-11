package main

import "fmt"

func main() {
	ages := map[string]int{
		"me":  2,
		"you": 3,
	}

	var data string
	fmt.Println("Enter your name")
	fmt.Scanln(&data)
	_, ok := ages[data]
	if !ok {
		fmt.Printf("%s\t is non existent\n", data)
	} else {
		fmt.Printf("%s\t your age is %d\n", data, ages[data])
	}
}

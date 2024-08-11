package main

import "fmt"

func main() {
	ages := map[string]int{
		"Mensur": 23,
		"Amira":  22,
		"Imran":  16,
		"Nejma":  9,
	}

	for k, v := range ages {
		fmt.Printf("%s\t%d\n", k, v)
	}

	var s string
	fmt.Println("Enter the key")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if age, ok := ages[s]; !ok {
		fmt.Printf("%s\n", s)
	} else {
		fmt.Printf("%v\n", age)
	}
}

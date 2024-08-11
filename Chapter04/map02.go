package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Mensur": 2,
		"Imran":  3,
		"Amira":  4,
		"Nejma":  5,
	}

	var names []string

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%v\n", name, ages[name])
	}

}

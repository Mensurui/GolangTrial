package Chapter02

import (
	"fmt"
	"strconv"
)

func trial() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("error")
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error")
	}

	pound := Pound(value)

	result := PoundToKilo(pound)

	fmt.Printf("%.2f of pound is %.2f of kilo", pound, result)
}

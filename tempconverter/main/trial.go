package main

import (
	"fmt"
	tempconverter "tempconverter/tempconver"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconverter.AbsoluteZeroC)
	fmt.Printf("Brrrr! in Kelvin%v\n", tempconverter.CToK(tempconverter.AbsoluteZeroC))

}

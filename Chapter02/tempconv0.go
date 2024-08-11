package chapter02

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	result := Fahrenheit(c*9/5 + 32)
	fmt.Println(result)
	return result
}
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

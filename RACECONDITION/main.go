package main

import (
	"example.com/raceConditions/condition1"
	"fmt"
)

func main() {
	//Alice:
	go func() {
		condition1.Deposite(200)               //A1
		fmt.Println("=", condition1.Balance()) //A2
	}()

	//Bob:
	go condition1.Deposite(100) //B
}

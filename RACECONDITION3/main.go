package main

import (
	"example.com/raceCondition3/condition3"
	"fmt"
	"time"
)

func main() {
	go func() {
		condition3.Deposit(300)
		fmt.Println("Alice Balance: ", condition3.Balance())
		condition3.Withdraw(200)
		fmt.Println("Alice Balance after withdrawing 200: ", condition3.Balance())
	}()

	go func() {
		condition3.Deposit(100)
		fmt.Println("Jake Balance: ", condition3.Balance())
		fmt.Println("Jake Balance: ", condition3.Balance())
		fmt.Println("Jake Balance: ", condition3.Balance())
		fmt.Println("Jake Balance: ", condition3.Balance())
	}()

	time.Sleep(2 * time.Second)
}

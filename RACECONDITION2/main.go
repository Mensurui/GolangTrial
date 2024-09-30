package main

import (
	"example.com/raceCondition2/condition2"
	"fmt"
	"time"
)

func main() {
	//Alice
	go func() {
		condition2.Deposite(2)
		fmt.Println("Balance: ", condition2.Balance())
		condition2.Withdraw(233)
		fmt.Println("Balance: ", condition2.Balance())
	}()

	go func() {
		condition2.Deposite(300)
		fmt.Println("Balance: ", condition2.Balance())
	}()

	time.Sleep(1 * time.Second)
}

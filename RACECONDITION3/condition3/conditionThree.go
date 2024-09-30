package condition3

import (
	"sync"
)

var (
	mu      sync.Mutex   //guards the balance
	mr      sync.RWMutex //To make readings occur parallely
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func Balance() int {
	mr.RLock()
	defer mr.RUnlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}

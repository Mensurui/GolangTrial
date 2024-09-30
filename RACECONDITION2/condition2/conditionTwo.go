package condition2

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)

func Deposite(amount int) { deposits <- amount }

func Withdraw(wamount int) { withdraws <- wamount }

func Balance() int { return <-balances }

func teller() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case wamount := <-withdraws:
			balance -= wamount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

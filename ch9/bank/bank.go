package bank

// Withdrawl represents a bank withdrawl
type Withdrawl struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdrawls = make(chan Withdrawl)

// Deposit money into the bank
func Deposit(amount int) { deposits <- amount }

// Balance returns the total balance
func Balance() int { return <-balances }

// Withdraw a specified amount
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawls <- Withdrawl{amount, ch}
	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount

		case w := <-withdrawls:
			if w.amount > balance {
				w.success <- false
				continue
			}
			balance -= w.amount
			w.success <- true

		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

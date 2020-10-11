package pointers_and_errors

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	fmt.Printf("address of balance in Balance is %v \n", &w.balance)
	return w.balance
}

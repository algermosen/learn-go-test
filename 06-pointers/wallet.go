package pointers

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

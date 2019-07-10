package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is a typealias to int
type Bitcoin int

// String is a function that returns a formatted string to describe a Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Stringer is an interface for types that implement String()
type Stringer interface {
	String() string
}

// Wallet is a struct that holds information about a wallet
type Wallet struct {
	balance Bitcoin
}

// Balance is a function that returns the `Bitcoin` balance of a `Wallet`
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit is a function that will add `Bitcoin`s to a `Wallet`
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds is an error describing that there are insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw is a function that removes balance from a `Wallet`. An error is thrown if the withdrawal
// amount is more than the avaiable balance.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

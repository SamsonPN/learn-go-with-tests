package pointerserrors

import (
	"errors"
	"fmt"
)

// int would've been fine for Wallet but
// we can create a Named type called bitcoin

// this is more descriptive than int
// and we can declare methods on it
type Bitcoin int

// we are implementing the Stringer interface on Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// lowercase field, method, func, etc are PRIVATE outside the package it's defined in!
	// only accessible inside the same package directly
	balance Bitcoin
}

// when you call  a function or a method, the arguments are COPIED
// func (w Wallet) Deposit(amount int) {

// by using a pointer to the Wallet, you are interacting with the actual argument
// these pointers to structs are called struct pointers
// these struct pointers are AUTOMATICALLY DEREFERENCED
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// Balance would still work correctly without the struct pointer since the copy
// would still have the same balance and we are not manipulating it
// but it is convention to keep the method receiver types the same for consistency
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// errors are values
// by turning assigning it to a var,
// if anyone changes the error message, it doesn't break the test
// this is a single source of truth for it
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// must import the "errors" package to use this
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

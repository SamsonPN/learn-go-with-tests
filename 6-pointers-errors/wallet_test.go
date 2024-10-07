package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// %p prints memory address in base 16
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	// after installing errcheck
	// `go install github.com/kisielk/errcheck@latest`
	// we ran `errcheck .` inside the directory where our tests are

	// this essentially tells us to check for different scenarios for that error
	// i.e. another test case to think about in this case
	t.Run("withdraw with funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		// since Bitcoin implements a Stringer method
		// we can now use %s for got and want
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	// nil is similar to null
	// errors can be nil b/c they are an interface
	// if a function takes arguments or returns values that are interfaces
	// they can be nillable
	if got == nil {
		// this will stop the test if it is called
		// don't want to make any more asserts on error returned if there isn't one
		t.Fatal("didn't get an error but wanted one")
	}

	// the error can be converted into a string using .Error() method
	// that's why we are able to compare it with want, which is a string
	// if got.Error != want (String at the time before refactoring)

	// so why are we doing this check?
	// it's to assert that we are returning the RIGHT error message to the user
	// before refactoring, we returned "oh no", which isn't very useful
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

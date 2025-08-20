package model

import (
	"testing"
)

func TestWallet(t *testing.T) {
	printBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}

	}
	printError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanter an error but didn't get one")
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		printBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		balance := Bitcoin(20)
		wallet := Wallet{balance: balance}
		err := wallet.Withdraw(Bitcoin(100))
		printError(t, err)
		printBalance(t, wallet, balance)
	})
}

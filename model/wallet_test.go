package model

import (
	"testing"
)

func printBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func printError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanter an error but didn't get one")
	}
	if want != nil {
		t.Errorf("got %s, but wanted %s", got, want)
	}
}

func printNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got and error but didn't want one.")
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		printBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		balance := Bitcoin(20)
		wallet := Wallet{balance: balance}
		err := wallet.Withdraw(Bitcoin(100))
		printNoError(t, err)
		printBalance(t, wallet, balance)
	})
}

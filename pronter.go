package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("测试比特币钱包", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()

		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

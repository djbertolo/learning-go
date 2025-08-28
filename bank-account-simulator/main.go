package main

import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
}

func main() {
	bankAccount := Account{
		balance: 1000,
	}

	bankAccount.Deposit(250)
	fmt.Println(bankAccount.balance)
	fmt.Println(bankAccount.Withdraw(300))
	fmt.Println(bankAccount.balance)
	fmt.Println(bankAccount.Withdraw(1500))
	fmt.Println(bankAccount.balance)
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("attempting to withdraw amount greater than balance")
	} else {
		a.balance -= amount
	}
	return nil
}

// Create a struct Account with fields:
// holderName (string)
// balance (float64)
// Now implement methods:
// Deposit(amount float64) → adds money to the balance.
// Withdraw(amount float64) → subtracts money (but only if enough balance exists).
// Info() → prints account holder name and current balance.
package main

import "fmt"

type Account struct {
	holderName string
	balance    float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount >= a.balance {
		fmt.Println("transaction inappropriate")
		return
	}
	a.balance -= amount
	fmt.Println("heres your withdrwal amount", amount)
}

func (a Account) Info() {
	fmt.Println("account holder is:-", a.holderName)
	fmt.Println("account's current balance is:-", a.balance)
}

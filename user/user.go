package user

import "fmt"

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(deposit float64) {
	u.Balance += deposit
}

func (u *User) Withdraw(amount float64) error {
	if u.Balance < amount {
		return fmt.Errorf("недостаточно средств")
	}
	u.Balance -= amount
	return nil
}

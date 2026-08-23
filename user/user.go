package user

import (
	"fmt"
	"sync"
)

type User struct {
	ID      string
	Name    string
	Balance float64
	mu      sync.Mutex
}

func (u *User) Deposit(deposit float64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Balance += deposit
}

func (u *User) Withdraw(amount float64) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	u.Balance -= amount
	return nil
}

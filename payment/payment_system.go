package payment

import (
	"app-money/transaction"
	"app-money/user"
	"fmt"
	"sync"
)

type PaymentSystem struct {
	Users        map[string]*user.User
	Transactions []transaction.Transcation
}

func (ps *PaymentSystem) Worker(ch <-chan transaction.Transcation, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range ch {
		err := ps.ProcessingTransactions(t)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (ps *PaymentSystem) AddUser(user *user.User) {
	ps.Users[user.ID] = user
}

func (ps *PaymentSystem) AddTransaction(transaction transaction.Transcation) {
	ps.Transactions = append(ps.Transactions, transaction)
}

func (ps *PaymentSystem) ProcessingTransactions(transaction transaction.Transcation) error {
	fromUser, ok := ps.Users[transaction.FromID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	toUser, ok := ps.Users[transaction.ToID]

	if !ok {
		return fmt.Errorf("user not found")

	}

	err := fromUser.Withdraw(transaction.Amount)
	if err != nil {
		return err
	}
	toUser.Deposit(transaction.Amount)
	return nil
}

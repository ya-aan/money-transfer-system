package payment

import (
	"app-money/transaction"
	"app-money/user"
	"fmt"
)

type PaymentSystem struct {
	Users        map[string]*user.User
	Transactions []transaction.Transcation
}

func (p *PaymentSystem) AddUser(user *user.User) {
	p.Users[user.ID] = user
}

func (p *PaymentSystem) AddTransaction(transaction transaction.Transcation) {
	p.Transactions = append(p.Transactions, transaction)
}

func (p *PaymentSystem) ProcessingTransactions(transaction transaction.Transcation) {
	fromUser, ok := p.Users[transaction.FromID]
	if !ok {
		fmt.Println("пользователь с данным ID не найден.")
		return
	}
	toUser, ok := p.Users[transaction.ToID]

	if !ok {
		fmt.Println("пользователь с данным ID не найден.")
		return
	}

	err := fromUser.Withdraw(transaction.Amount)
	if err != nil {
		fmt.Println(err)
		return
	}
	toUser.Deposit(transaction.Amount)
}

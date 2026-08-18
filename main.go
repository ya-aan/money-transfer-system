package main

import (
	"app-money/payment"
	"app-money/transaction"
	"app-money/user"
	"fmt"
)

func main() {
	user1 := &user.User{
		ID:      "1",
		Name:    "Yasha",
		Balance: 200,
	}
	user2 := &user.User{
		ID:      "2",
		Name:    "Sasha",
		Balance: 50,
	}

	paymentSystem := &payment.PaymentSystem{
		Users: make(map[string]*user.User),
	}

	paymentSystem.AddUser(user1)
	paymentSystem.AddUser(user2)
	transaction1 := transaction.Transcation{
		FromID: "1",
		ToID:   "2",
		Amount: 100,
	}
	paymentSystem.AddTransaction(transaction1)
	for _, transaction := range paymentSystem.Transactions {
		paymentSystem.ProcessingTransactions(transaction)
	}
	fmt.Println("Баланс первого пользователя:", user1.Balance)
	fmt.Println("Баланс второго пользователя:", user2.Balance)
}

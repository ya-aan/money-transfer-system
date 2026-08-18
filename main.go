package main

import (
	"app-money/payment"
	"app-money/transaction"
	"app-money/user"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
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
		Amount: 10,
	}
	transaction2 := transaction.Transcation{
		FromID: "1",
		ToID:   "2",
		Amount: 10,
	}

	paymentSystem.AddTransaction(transaction1)
	paymentSystem.AddTransaction(transaction2)
	ch := make(chan transaction.Transcation, len(paymentSystem.Transactions))
	for i := 0; i < len(paymentSystem.Transactions); i++ {
		wg.Add(1)
		go paymentSystem.Worker(ch, &wg)
	}
	for _, t := range paymentSystem.Transactions {
		ch <- t
	}

	close(ch)
	wg.Wait()
	fmt.Println("Баланс первого пользователя:", user1.Balance)
	fmt.Println("Баланс второго пользователя:", user2.Balance)
}

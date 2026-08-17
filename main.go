package main

import (
	"app-money/user"
	"fmt"
)

func main() {
	// TODO: Создайте несколько пользователей user1, user2
	user1 := &user.User{
		ID:      "1",
		Name:    "Yasha",
		Balance: 200,
	}
	user2 := &user.User{
		ID:      "2",
		Name:    "Sasha",
		Balance: 5052,
	}
	// TODO: Проверьте работу методов Deposit и Withdraw
	user1.Deposit(50)
	user1.Withdraw(200)
	user2.Deposit(555)
	user2.Withdraw(4500)
	// TODO: Выведите информацию о пользователях
	fmt.Println(user1)
	fmt.Println(user2)

}

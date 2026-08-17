package user

import "fmt"

// TODO: Создайте структуру User с полями ID, Name и Balance
type User struct {
	ID      string
	Name    string
	Balance float64
}

// TODO: Реализуйте метод Deposit
func (u *User) Deposit(deposit float64) {
	u.Balance += deposit
}

// TODO: Реализуйте метод Withdraw с проверкой на недостаток средств
func (u *User) Withdraw(amount float64) {
	if u.Balance < amount {
		fmt.Println("Денег нет")
		return
	}
	u.Balance -= amount
}

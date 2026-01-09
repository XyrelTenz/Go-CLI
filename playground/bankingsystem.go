// Package playground
package playground

import (
	"fmt"
)

type Accounts struct {
	Pin           int
	AccountNumber int
	Balance       int
}

type Pins interface {
	Auth(pin int, account int)
}

type Util interface {
	Authentication(Pin int, AccountNumber int)
}

func Withdraw(b *Accounts) {

	fmt.Println("Withdraw")

}

func Deposit(b *Accounts) {

	fmt.Println("Deposit")
}

func CheckBalance(b *Accounts) {

	fmt.Printf("Your Balance is: %d \n", b.Balance)
}

var (
	pin           int
	accountNumber int
)

func Authentication(method Util, Pin int, AccountNumber int) {

	method.Authentication(Pin, AccountNumber)

}

func Login() {

	accounts := &Accounts{
		Pin:           1234,
		AccountNumber: 123123123,
	}

	fmt.Print("Enter Pin: ")

	_, error := fmt.Scanln(&pin)

	if error != nil {

		fmt.Println("Invalid Data Types")
		panic(error)
	}

	fmt.Print("Enter Account Number: ")
	_, err := fmt.Scanln(&accountNumber)

	if err != nil {

		fmt.Println("Invalid Data Types")
		panic(err)

	}

	Authentication(accounts, pin, accountNumber)

}

func (a *Accounts) Authentication(pin int, account int) {

	if a.AccountNumber != account && a.Pin != pin {

		fmt.Println("Invalid Credentials")

	} else {
		BankingSystem()

	}

}

func BankingSystem() {

	Accounts := &Accounts{
		Balance: 0,
	}

	for {

		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Balance")
		fmt.Println("4. Logout")

		var operations int

		fmt.Print("Choose Operations: ")
		_, error := fmt.Scanln(&operations)

		if error != nil {
			fmt.Println("Invalid Data types")
			panic(error)
		}

		switch operations {
		case 1:
			Deposit(Accounts)
		case 2:
			Withdraw(Accounts)
		case 3:
			CheckBalance(Accounts)
		case 4:
			fmt.Println("Logout...")
			return
		default:
			fmt.Println("Invalid Operations")

		}

	}

}

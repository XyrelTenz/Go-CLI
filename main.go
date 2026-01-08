package main

import (
	"bufio"
	"fmt"
	"os"
)

//TODO: 1. Learn About Garoutines
//TODO: 2. Learn How to Manipulate Function and Struct
//TODO: 3. Learn Struct types

type Utils interface {
	Pay(amount int)
}

type Gcash struct {
	PhoneNumber int
}

type Maya struct {
	AccountNumber int
}

type Cash struct{}

func (g *Gcash) Pay(amount int) {

	fmt.Println("Paid with Gcash")
	fmt.Printf("Phone Number: %d\n", g.PhoneNumber)
	fmt.Printf("Total: %d \n", amount)

}

func (m *Maya) Pay(amount int) {

	fmt.Println("Paid with Maya")
	fmt.Printf("Account Number: %d\n", m.AccountNumber)
	fmt.Printf("Total: %d \n", amount)

}

func (c *Cash) Pay(amount int) {

	fmt.Println("Paid with Cash")
	fmt.Printf("Total: %d \n", amount)

}

func PaymentProcess(method Utils, amount int) {

	method.Pay(amount)
}

func Fibonacci() func() int {

	a, b := 0, 1

	return func() int {

		current := a

		a, b = b, a+b

		return current

	}

}

func main() {
	// pointers.BasicPointer()
	// p.Hero()

	maya := &Maya{

		AccountNumber: 400,
	}

	gcash := &Gcash{

		PhoneNumber: 12345678910,
	}

	cash := &Cash{}

	PaymentProcess(maya, 100)
	PaymentProcess(gcash, 100)
	PaymentProcess(cash, 100)

	nextFib := Fibonacci()

	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())

}

type Person struct {
	StudentName string
	StudentAge  int
}

var (
	name string = "Xyrel"
	age  int    = 18
)

// OS and Bufio

func StudentInfo() {
	age := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name: ")

	// Comma Syntax only Contains OK and Err
	input, _ := age.ReadString('\n')

	fmt.Print(input)
	fmt.Printf("Type is %T  \n", input)

}

// Conversions

// Login

func LoginSystem() func() {

	username := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")

	usernameInput, _ := username.ReadString('\n')

	password := bufio.NewReader(os.Stdin)
	fmt.Print("Password: ")

	passwordInput, _ := password.ReadString('\n')

	auth := Authentication(usernameInput, passwordInput)

	return auth

}

// Authentication Here
func Authentication(username string, password string) func() {
	if username != "xyrel" && password != "123" {
		fmt.Print("Invalid Credentials")
	} else {
		fmt.Print("Authentication Successs")
	}

	return BankingSystem

}

func BankingSystem() {
	fmt.Print("This is Banking System")
}

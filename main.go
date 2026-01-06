package main

import (
	"bufio"
	"fmt"
	"os"

	p "golangapp/playground"
	"golangapp/pointers"
)

//TODO: 1. Learn About Pointers and References
//TODO: 2. Learn About Arrays and Maps
//TODO: 3. Learn About Garoutines
//TODO: 4. Learn How to Manipulate Function and Struct
//TODO: 5. Learn Struct types

func main() {
	// Variable and Declarations
	var name string = "Xyrel"
	// Shorthand
	name2 := "Xyrel2"

	// Types Int and Value
	var age uint8 = 255
	var age2 uint32 = 1000
	var age3 uint64 = 200000

	// Types Float and Value
	var floatAge float32 = 10.99999
	var floatAge2 float64 = 10.999999

	// Boolean Expression
	isAdult := false
	isMinor := true

	fmt.Println(name)
	fmt.Println(name2)

	fmt.Println(age)
	fmt.Println(age2)
	fmt.Println(age3)

	fmt.Println(floatAge)
	fmt.Println(floatAge2)

	fmt.Println(isAdult)
	fmt.Println(isMinor)

	// Student Info
	// StudentInfo()

	// BankingSystem
	// LoginSystem()

	// Pointer and References
	ages := 19

	ptr := &ages

	fmt.Printf("Value: %v \n", *ptr)

	// Array
	// ... meaning can accept of any length of an Array
	fruits := [...]string{
		"Apple",
		"Mango",
		"Papaya",
	}

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	fmt.Print(name)

	StudentInfo := &Person{
		StudentName: "Xyrel",
		StudentAge:  19,
	}

	fmt.Println(StudentInfo)

	pointers.BasicPointer()
	p.Hero()

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

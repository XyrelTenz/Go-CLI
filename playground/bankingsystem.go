// Package playground
package playground

import (
	"bufio"
	"fmt"
	"os"
)

//TODO: Refactor This and Apply What I've been learn such as Closure, Inteface, Pointers, Comma OK

var (
	name string = "Xyrel"
	age  int    = 18
)

func BankingSystem() {

	fmt.Println("1. Withdraw")
	fmt.Println("2. Balance")
	fmt.Print("3. Deposit")
	fmt.Print("4. Exit")

	ReadOperations := bufio.NewReader(os.Stdin)

	operations, err := ReadOperations.ReadByte()

	if err != nil {
		fmt.Print(err)
	}

	switch operations {
	case 1:
		fmt.Print("Withdraw")
	case 2:
		fmt.Print("Balance")
	case 3:
		fmt.Print("Deposit")
	case 4:
		fmt.Print("Exit")
	default:
		fmt.Print("Invalid Operations")

	}
}

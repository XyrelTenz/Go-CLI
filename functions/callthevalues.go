// Package functions
package functions

import "fmt"

func Callers(num int) {

	num = 100

}

func CallTheValue() {

	num := 10

	newNum := &num

	Callers(*newNum)
	fmt.Println(*newNum)

}

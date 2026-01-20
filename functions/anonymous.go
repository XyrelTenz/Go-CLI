// Package functions
package functions

import "fmt"

func Anonymous() {

	go func(name string) {

		fmt.Println(name)

	}("Xyrel")

	// Assign to an variables

	greetings := func(name string) {
		fmt.Println(name)

	}

	greetings("Hello")

}

// Package pointers demonstrates how memory addresses work in Go.
package pointers

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Scenario A: Pass by Value (The "Photocopy")
// This function gets a COPY. Changing it does nothing to the original.
func changeNameByValue(p Person) {
	p.Name = "Alice"
	fmt.Printf("Inside Value Function: %s\n", p.Name)
}

// This function gets the ADDRESS. Changing it updates the original.
func changeNameByPointer(p *Person) {
	p.Name = "Bob"
	fmt.Printf("Inside Pointer Function: %s\n", p.Name)
}

func BasicPointer() {
	// Initialize our person
	// We can use & right here to make 'me' a pointer immediately if we want,
	// but let's start with a value to see the difference clearly.
	me := Person{Name: "Xyrel", Age: 19}

	fmt.Println("--- START ---")
	fmt.Printf("Original: %s\n", me.Name)

	changeNameByValue(me)
	fmt.Printf("After Value Call: %s (No Change!)\n", me.Name)

	fmt.Println("---")

	// Note: We use '&me' to send the address of 'me'
	changeNameByPointer(&me)
	fmt.Printf("After Pointer Call: %s (Changed!)\n", me.Name)
}

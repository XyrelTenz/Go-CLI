// Package concurrency
package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// TODO: 1. Understand What Time Package is
// TODO: 2. Understand What Sync Package is
// TODO: 3. Understand Anonymous Function and How to use it
// TODO: 4. Understand Garoutines
func RunConcurrency() {
	// WaitGroup acts as a counter for active threads
	var wg sync.WaitGroup

	fmt.Println("--- Starting Async Tasks ---")

	// Launch 3 Goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter

		// Anonymous function (Closure)
		go func(id int) {
			defer wg.Done() // Decrement counter when function finishes

			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(500 * time.Millisecond) // Simulate work
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	// Block here until counter is 0
	wg.Wait()
	fmt.Println("All workers finished.")
}

func Anonymous() {

	number := 5
	go func() {
		number++
		fmt.Println("Increase")

	}()

	fmt.Println(string(number))

}

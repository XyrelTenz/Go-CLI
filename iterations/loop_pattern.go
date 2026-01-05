// Package iteration
package iteration

import "fmt"

func RunPatterns() {
	data := []int{10, 20, 30, 40, 50, 60, 70}

	// 1. SLICING PATTERNS (Critical for Algorithms)
	// format: array[start_inclusive : end_exclusive]
	part := data[1:4] // Index 1, 2, 3 (Val: 20, 30, 40)
	fmt.Printf("Slice: %v\n", part)

	// 2. MANUAL INDEX MANIPULATION (The "Two Pointer" Technique)
	// Example: Reversing an array in-place
	fmt.Println("--- Two Pointer Reverse ---")
	left, right := 0, len(data)-1

	for left < right {
		// Swap values
		data[left], data[right] = data[right], data[left]

		// Move indices manually
		left++
		right--
	}
	fmt.Printf("Reversed: %v\n", data)

	// 3. ADVANCED LOOPS: "Labelled Breaks"
	// Used to break out of nested loops (like in matrix searches)
	fmt.Println("--- Labelled Break ---")
	matrix := [][]int{
		{1, 2},
		{3, 4}, // Let's search for 4
	}

SearchLoop: // Define a label
	for r, row := range matrix {
		for c, val := range row {
			if val == 4 {
				fmt.Printf("Found 4 at [%d][%d]\n", r, c)
				break SearchLoop // Breaks the OUTER loop immediately
			}
		}
	}
}

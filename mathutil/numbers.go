// Package mathutil
package mathutil

import (
	"fmt"
	"math"
)

// Go's math library works on float64 usually.
// For algorithms, we often need helpers for Ints.

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func DemoMath() {
	// Using Standard Math for Floats
	val := 16.0
	fmt.Printf("Sqrt of 16: %.2f\n", math.Sqrt(val))

	// Using our helpers for Int Algorithms
	fmt.Printf("Min of 10, 5: %d\n", Min(10, 5))
}

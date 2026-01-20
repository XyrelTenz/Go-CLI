// Package functions
package functions

// Accept Unlimited Numbers of Arguement
func Variadic(nums ...int) int {

	total := 0

	for _, n := range nums {

		total += n

	}

	return total

}

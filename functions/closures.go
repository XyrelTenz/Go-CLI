// Package functions

package functions

// Closures is just returning an function inside the function
func Closures() func() int {

	count := 0

	return func() int {

		count++
		return count

	}

}

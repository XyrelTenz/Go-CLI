package pointers

import "fmt"

// Hidden Pointers
func modifyMap(m map[string]int) {

	m["score"] = 100
}

func main() {
	stats := map[string]int{"score": 10}
	modifyMap(stats)
	fmt.Println(stats["score"]) // Changed without using *map
}

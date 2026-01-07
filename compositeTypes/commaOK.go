package compositetypes

import "fmt"

func CommaOK() {
	scores := map[string]int{
		"Alice": 95,
		"Bob":   82,
	}

	score, ok := scores["Alice"]
	if ok {
		fmt.Printf("Alice's score is %d\n", score)
	}

	score, ok = scores["Charlie"]
	if !ok {
		fmt.Println("Charlie was not found")
	}
}

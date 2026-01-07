package compositetypes

import "fmt"

//TODO: Learn about usecase of CommaOK

func KeyChecker() {
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

func typeAssertion() {

	var data interface{} = "Hello World"

	str, ok := data.(string)

	if ok {

		fmt.Printf("%s is a string", str)
	}

	int, ok := data.(int)

	if !ok {

		fmt.Println("Is not an Integeter")
	} else {
		fmt.Println(int)
	}

}

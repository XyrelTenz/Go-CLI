package pointers

import "fmt"

type Structure struct {
	value int
}

var testStruct = Structure{
	value: 0,
}

func AddNumbers(x, y int) int {

	return x + y

}
func MemoryManagement() {

	t := &Structure{}

	go func() {
		fmt.Println(AddNumbers(5, t.value))
	}()

}

package iteration

import (
	"fmt"
	"sort"
)

func MapIteration() {

	info := map[int]string{

		3: "ID: ",
		2: "Name: ",
		1: "Gender: ",
	}

	keys := make([]int, 0, len(info))

	for key := range info {

		keys = append(keys, key)

	}

	sort.Ints(keys)

	for _, v := range keys {

		fmt.Printf("Key: %d, Value: %s\n", v, info[v])

	}

}

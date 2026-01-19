package concurrency

import "fmt"

func sum(s []int, c chan int) {

	sum := 0

	for _, v := range s {

		sum += v
	}

	c <- sum

}

func Channels() {

	s := []int{
		7, 2, 3, 0, -8, 10,
	}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	animal := make(chan string)
	human := make(chan string)

	go func() {

		msg := <-animal
		fmt.Println("Animals: ", msg)

	}()

	go func() {

		msg := <-human
		fmt.Println("Human: ", msg)

	}()

	behavior := []string{

		"Fly",
		"Run",
	}

	for _, v := range behavior {

		if v == "Fly" {

			animal <- "Animals is Flying"

		}

		if v == "Run" {

			human <- "Human is Running"

		}

	}

	fmt.Scanln()

}

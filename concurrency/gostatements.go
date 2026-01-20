package concurrency

import (
	"fmt"
	"time"
)

func GoStatement() {

	message := make(chan string)

	go func(ch chan<- string) {

		time.Sleep(1 * time.Millisecond)

		message <- "Hello World"

	}(message)

	msg := <-message

	fmt.Println(msg)
}

func array() {

}

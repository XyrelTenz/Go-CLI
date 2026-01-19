package playground

import (
	"fmt"
	"net/http"
	"time"
)

func WebsiteHealth() {

	start := time.Now()
	c := make(chan string)

	website := []string{

		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, v := range website {

		CheckURL(v, c)

	}

	for v := range len(website) {

		fmt.Println(<-c, v)

	}

	fmt.Printf("Time since started: %s", time.Since(start))

}

func CheckURL(url string, c chan string) {

	_, err := http.Get(url)

	if err != nil {

		c <- url + " is Down"

	}

	c <- url + " is Up"
}

package playground

import (
	"fmt"
	"time"
)

func Time(song []string, t time.Duration) {

	for _, v := range song {

		time.Sleep(t)
		fmt.Println(v)

	}

}

func Lyrics() {

	song := []string{

		"Hello",
		"HAHA",
	}

	Time(song, 1*time.Second)

}

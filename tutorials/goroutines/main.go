package main

import (
	"fmt"
	"time"
)

func f(from string, to ...int) {
	tov := 3

	if len(to) > 0 {
		tov = to[0]
	}

	for i := 0; i < tov; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine", 50)
	go f("goroutine2", 100)

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

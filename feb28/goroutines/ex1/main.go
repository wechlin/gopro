package main

import (
	"fmt"
)

func greet(n int, finished chan<- bool) {
	fmt.Println("Hello World", n)
	// Send to channel
	finished <- true
}

func main() {

	var times = 5
	done := make(chan bool)

	for n := 0; n < times; n++ {
		go greet(n, done)
	}

	for n := 0; n < times; n++ {
		// Receive from channel
		<-done
	}

}

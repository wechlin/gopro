package main

import (
	"fmt"
	"time"
)

func main() {
	// Create timeout objects that will send the current time on the .C field once it finishes
	timeout := time.NewTimer(3 * time.Second)

	for {
		// Select works like select(2) or poll(2) syscall, but on channels to read from or
		// write to, and will pick whichever is available (and use the default case if none are ready)
		select {
		case tm := <-timeout.C:
			fmt.Println("Current time is", tm)
			return
		//Default case is optional
		default:
			fmt.Println("Waiting...")
		}
	}
}

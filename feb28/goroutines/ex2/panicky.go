package main

import "fmt"

func main() {
	blocker := make(chan bool)

	fmt.Println(<-blocker)
}

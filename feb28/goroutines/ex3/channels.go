package main

import "fmt"

type object struct {
	name string
}

func main() {
	norman := object{"Norman"}

	valueChannel := make(chan object)
	done := make(chan bool)

	go func(done chan<- bool) {
		local := <-valueChannel
		local.name = "Norbert"
		fmt.Println(local)
		done <- true
	}(done)

	valueChannel <- norman
	<-done
	fmt.Println(norman)

	refChannel := make(chan *object)

	go func(done chan<- bool) {
		local := <-refChannel
		local.name = "Nyman"
		fmt.Println(local)
		done <- true
	}(done)

	refChannel <- &norman
	<-done
	fmt.Println(norman)
}

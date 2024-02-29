package main

import "fmt"

func filter(p int, input <-chan int, output chan<- int) {
	for n := range input {
		if n%p != 0 {
			output <- n
		}
	}
}

func main() {
	head := make(chan int)

	go func() {
		river := head
		for n := 2; ; n++ {
			river <- n
		}
	}()

	for {
		p := <-head
		fmt.Println(p)
		tail := make(chan int)
		go filter(p, head, tail)

		head = tail
	}
}

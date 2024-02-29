package main

import "fmt"

func main() {
	target := "World"

	go func() {
		target = "Planet"
	}()

	fmt.Println("Hello", target)
}

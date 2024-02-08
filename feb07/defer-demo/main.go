package main

import (
	"fmt"
	"os"
)

func hello() {
	defer fmt.Println("[leaving now]")

	fmt.Println("Hello World!")
}

func main() {

	f, err := os.Create("file.txt")
	if err != nil {
		// log.Fatal would be better
		fmt.Println("Unable to create file")
		return
	}
	defer f.Close()

	hello()
}

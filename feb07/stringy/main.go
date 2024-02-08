package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <number>\n", os.Args[0])
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to convert '%s': %s", os.Args[1], err)
	}

	fmt.Println(num * 2)
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <word>\n", os.Args[0])
	}

	fh, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Unable to open 'output.txt'", err)
	}
	defer fh.Close()

	fmt.Fprintf(fh, "The word is '%s'\n", os.Args[1])
}

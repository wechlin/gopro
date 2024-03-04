package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage %s <pattern> <path>\n", os.Args[0])
	}
}

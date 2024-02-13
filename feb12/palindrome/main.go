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

	word := os.Args[1]

	for n := 0; n < len(word)/2; n++ {
		// Only works for ASCII text, need to import "unicode/utf8" package
		// or convert the word to a []rune
		if word[len(word)-1-n] != word[n] {
			fmt.Println("Not a palindrome")
			return
		}
	}

	fmt.Println("Is a palindrome")
}

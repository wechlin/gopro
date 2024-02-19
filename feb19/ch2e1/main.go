package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func MatchPrefixes(sentence string, startWith rune) []string {
	words := strings.Fields(sentence)

	results := make([]string, 0)

	for _, word := range words {
		if strings.HasPrefix(word, string(startWith)) {
			log.Println(cap(results))
			results = append(results, word)
		}
	}
	log.Println(cap(results))

	return results
}

func main() {
	for _, found := range MatchPrefixes(os.Args[1], 'e') {
		fmt.Println(found)
	}
}

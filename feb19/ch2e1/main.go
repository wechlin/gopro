package main

import (
	"fmt"
	"os"
	"strings"
)

func MatchPrefixes(sentence string, startWith rune) []string {
	words := strings.Fields(sentence)

	results := make([]string, 0)

	for _, word := range words {
		if strings.HasPrefix(word, string(startWith)) {
			results = append(results, word)
		}
	}

	return results
}

func main() {
	for _, found := range MatchPrefixes(os.Args[1], 'e') {
		fmt.Println(found)
	}
}

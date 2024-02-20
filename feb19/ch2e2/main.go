package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	frequency := make(map[rune]int)

	for _, arg := range os.Args[1:] {
		for _, letter := range arg {
			frequency[letter] += 1
		}
	}

	type letterCount struct {
		letter rune
		count  int
	}

	results := make([]letterCount, 0)

	for letter, count := range frequency {
		results = append(results, letterCount{letter, count})
	}

	slices.SortFunc(results, func(a, b letterCount) int {
		return int(a.letter - b.letter)
	})

	for _, item := range results {
		fmt.Println(string(item.letter), item.count)
	}
}

package main

import (
	"fmt"
	"os"
	"sort"
)

type letterCount struct {
	letter rune
	count  int
}

type freqResult []letterCount

func (fr *freqResult) Len() int {
	return len(*fr)
}

func (fr *freqResult) Less(i, j int) bool {
	return (*fr)[i].letter < (*fr)[j].letter
}

func (fr *freqResult) Swap(i, j int) {
	(*fr)[i], (*fr)[j] = (*fr)[j], (*fr)[i]
}

func main() {
	frequency := make(map[rune]int)

	for _, arg := range os.Args[1:] {
		for _, letter := range arg {
			frequency[letter] += 1
		}
	}

	results := freqResult{}

	for letter, count := range frequency {
		results = append(results, letterCount{letter, count})
	}

	sort.Sort(&results)

	for _, item := range results {
		fmt.Println(string(item.letter), item.count)
	}
}

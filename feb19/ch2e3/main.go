package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

func processFile(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func main() {
	count := flag.Int("count", -1, "print first 'n' results")
	unique := flag.Bool("unique", false, "no duplicates")

	flag.Parse()

	results := make([]string, 0)

	for _, fname := range flag.Args() {
		words, err := processFile(fname)
		if err != nil {
			log.Printf("Unable to read words from '%s': %v", fname, err)
			continue
		}
		results = append(results, words...)
	}

	slices.Sort(results)

	var lastWord string
	for _, word := range results {
		if *unique && lastWord == word {
			continue
		}

		if *count == 0 {
			break
		}
		*count--

		fmt.Println(word)
		lastWord = word
	}
}

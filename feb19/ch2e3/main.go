package main

import (
	"bufio"
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
	//count := flag.Int("count", -1, "print first 'n' results")
	//unique := flag.Bool("unique", false, "no duplicates")

	results := make([]string, 0)

	for _, fname := range os.Args[1:] {
		words, err := processFile(fname)
		if err != nil {
			log.Printf("Unable to read words from '%s': %v", fname, err)
			continue
		}
		results = append(results, words...)
	}

	slices.Sort(results)

	for _, word := range results {
		fmt.Println(word)
	}
}

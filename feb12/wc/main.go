package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	totalLines, totalWords, totalBytes := 0, 0, 0

	for _, fname := range os.Args[1:] {
		lines, words, bytes := countFile(fname)
		fmt.Printf("  %d   %d   %d  %s\n", lines, words, bytes, fname)
		totalLines += lines
		totalWords += words
		totalBytes += bytes
	}
	fmt.Printf("  %d   %d   %d  total\n", totalLines, totalWords, totalBytes)
}

func countFile(filename string) (lines, words, bytes int) {
	fh, err := os.Open(filename)
	if err != nil {
		log.Printf("Unable to open '%s': %v", filename, err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		line := scanner.Text()
		lines++

		words += len(strings.Fields(line))

		// DISLIKE, rewrite to not be an arbitrary +1 for newline
		bytes += len(line) + 1
	}

	return
}

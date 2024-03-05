package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"runtime/pprof"
	"strings"
	"sync"
)

var wc sync.WaitGroup

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage %s <text> <path>\n", os.Args[0])
	}
	pattern := os.Args[1]
	root := os.Args[2]
	//TODO vet that `root` is a directory
	// After running, invoke `go tool pprof ch6e2 cpuprofile`
	profileOutput, _ := os.Create("cpuprofile")

	pprof.StartCPUProfile(profileOutput)

	fs.WalkDir(os.DirFS(root), ".", func(fname string, d fs.DirEntry, problem error) error {
		if d.IsDir() {
			return nil
		}

		wc.Add(1)
		go searchFile(path.Join(root, fname), pattern)

		return nil
	})

	wc.Wait()
	pprof.StopCPUProfile()
}

func searchFile(fname, needle string) {
	defer wc.Done()
	fh, err := os.Open(fname)
	if err != nil {
		log.Printf("Unable to open %s: %v", fname, err)
		return
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		haystack := scanner.Text()
		if strings.Contains(haystack, needle) {
			fmt.Println(fname)
			return
		}
	}

	if err = scanner.Err(); err != nil {
		log.Printf("Error reading %s: %v", fname, err)
	}
}

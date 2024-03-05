package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage %s <filename> <path>\n", os.Args[0])
	}
	needle := os.Args[1]
	root := os.Args[2]
	//TODO vet that `root` is a directory
	// After running, invoke `go tool pprof ch6e2 cpuprofile`

	found := make(chan string)

	go func() {
		fs.WalkDir(os.DirFS(root), ".", func(fname string, d fs.DirEntry, problem error) error {
			if d.Name() == needle {
				found <- fname
			}

			return nil
		})
		close(found)
	}()

	entry, ok := <-found
	if ok {
		fmt.Println(entry)
	} else {
		fmt.Println("No matches found")
	}

}

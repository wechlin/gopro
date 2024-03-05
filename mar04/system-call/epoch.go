package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/unix"
)

func main() {
	tt, err := unix.Time(nil)
	if err != nil {
		log.Fatalf("Unable to retrieve time: %v", err)
	}

	fmt.Println(tt)
}

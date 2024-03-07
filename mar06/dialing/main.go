package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Run `nc -v -l localhost 8080` to start listening, then after this program connects to it, type text
// into nc which will show up in this program' output.

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	fmt.Println(reader.ReadString('\n'))
}

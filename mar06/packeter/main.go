package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp6", ":24601")
	if err != nil {
		log.Fatal(err)

	}

	defer conn.Close()

	buf := make([]byte, 1024)
	num, remote, err := conn.ReadFrom(buf)
	if err != nil {
		log.Print(err)
	}

	recvd := string(buf[:num])

	fmt.Printf("Received %q from %s\n", recvd, remote)
}

package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	srv, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		client, err := srv.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go func() {
			defer client.Close()

			w := bufio.NewWriter(client)

			w.WriteString(time.Now().String())
			// bufio made Liam sad :(
			// The .Flush() is needed to make sure the .Write() underneath happens
			// before the client is .Closed
			w.Flush()
		}()
	}
}

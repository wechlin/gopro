package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func getdemo() {

	page, err := http.Get("http://example.org")
	if err != nil {
		log.Fatal(err)
	}
	defer page.Body.Close()

	buf := bufio.NewReader(page.Body)

	fmt.Print(buf.ReadString('\n'))
}

func main() {
	getdemo()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Query())
		w.Write([]byte("Hello World"))
	})

	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("/home/heptadecagram/go/gopro/"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

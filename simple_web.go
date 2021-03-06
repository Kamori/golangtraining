package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	reqHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", reqHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

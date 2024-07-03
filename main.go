package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":80"
	fmt.Println("Listening @ port", port)
	http.HandleFunc("/", Handler)
	http.ListenAndServe(port, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

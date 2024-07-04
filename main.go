package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":80"
	fmt.Println("Listening @ port", port)
	http.HandleFunc("/", Handler)
	http.ListenAndServe(port, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	message := "Hello, World!"
	envVar, ok := os.LookupEnv("TEST_ENV_VAR")
	if ok {
		message = fmt.Sprintf("Hello, %s!", envVar)
	}
	fmt.Fprint(w, message)
}

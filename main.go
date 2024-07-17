package main

import (
	"fmt"
	"io"
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
	forwardingUrl, ok := os.LookupEnv("FORWARDING_URL")
	fmt.Println("Forwarding URL:", forwardingUrl)
	if ok {
		res, err := http.Get(forwardingUrl)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(w, "Error reading body: %v", err)
			return
		}

		w.Write(body)
	} else {
		message := "Hello, World!"
		envVar, ok := os.LookupEnv("TEST_ENV_VAR")
		if ok {
			message = fmt.Sprintf("Hello, %s!", envVar)
		}
		fmt.Fprint(w, message)
	}
}

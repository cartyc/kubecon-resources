package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Kubecon!")
}

func main() {
	http.HandleFunc("/", hello)

	port := ":8080"

	http.ListenAndServe(port, nil)
}

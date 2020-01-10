package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(hello)
	if err := http.ListenAndServe("localhost:8080", h); err != nil {
		log.Fatalf("error: :isteneing and serving %s", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Heello you requested with %s %s", r.Method, r.URL.Path)
}

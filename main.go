package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to your Go application!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
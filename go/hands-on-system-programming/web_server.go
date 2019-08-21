package main

import (
	"fmt"
	"net/http"
)

type customerHandler int

func (c *customerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", *c)
	*c++
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	mux.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "cat!")
	})

	mux.Handle("/bye", new(customerHandler))

	if err := http.ListenAndServe(":3000", mux); err != nil {
		panic(err)
	}
}

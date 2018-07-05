package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/",r)
	http.Handle("/2",r)

	http.ListenAndServe(":8080", nil)
}

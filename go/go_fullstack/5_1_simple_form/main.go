package main

import "github.com/gorilla/mux"
import (
	"5_1_simple_form/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	http.ListenAndServe(":2992", r)
}

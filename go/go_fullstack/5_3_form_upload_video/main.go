package main

import "github.com/gorilla/mux"
import (
	"net/http"
	"5_3_form_upload_video/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload-image", handlers.UploadVideoHandler).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":2992", r)
}

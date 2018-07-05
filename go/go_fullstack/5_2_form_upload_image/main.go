package main

import "github.com/gorilla/mux"
import (
	"github.com/5_2_form_upload_image/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload-image", handlers.UploadImageHandler).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe("192.168.1.51:2992", r)
}

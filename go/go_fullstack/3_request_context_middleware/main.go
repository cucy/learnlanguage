package main

import (
	ghandlers "github.com/gorilla/handlers"

	"3_request_context/handlers"
	"3_request_context/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)

	http.Handle("/",
		middleware.ContextExapmleHandler(
			middleware.PanicRecoveryHandler(
				ghandlers.LoggingHandler(os.Stdout, r))))

	http.ListenAndServe(":8081", nil)

}

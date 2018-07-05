package main

import (
	"4_rest_api/verdor/endpoints"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	// http://127.0.0.1:1991/restapi/socialmediapost/zhansan
	r.HandleFunc("/restapi/socialmediapost/{username}", endpoints.FetchPostsEndpoint).Methods("GET")

	r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.CreatePostEndpoint).Methods("POST")
	r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
	r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")
	http.ListenAndServe("127.0.0.1:1991", r)

}

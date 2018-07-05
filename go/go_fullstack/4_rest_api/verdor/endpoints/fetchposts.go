package endpoints

import (
	"github.com/gorilla/mux"
	"net/http"

	"4_rest_api/verdor/models/socialmedia"
	"encoding/json"
)

func FetchPostsEndpoint(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	if v["username"] == "zhansan" {
		mockPosts := make([]socialmedia.Post, 3)

		post1 := socialmedia.NewPost(
			"EngineerKamesh", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post2 := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["happy"], "Tour of Go", "Check out the Tour of Go!", "https://tour.golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post3 := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["hopeful"], "Go Playground", "Check out the Go Playground!", "https://playground.golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})

		mockPosts = append(mockPosts, *post1)
		mockPosts = append(mockPosts, *post2)
		mockPosts = append(mockPosts, *post3)
		json.NewEncoder(w).Encode(mockPosts)
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}


# 1
```

r.HandleFunc("/restapi/socialmediapost/{username}", endpoints.FetchPostsEndpoint).Methods("GET")
r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.CreatePostEndpoint).Methods("POST")
r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")
```

```
package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/EngineerKamesh/gofullstack/volume2/section2/gopherface/models/socialmedia"
)

func FetchPostsEndpoint(w http.ResponseWriter, r *http.Request) {

	// TODO: Implement fetching posts for a given user

	// We are going to create some mock data and send it out in JSON
	// format.

	// We will actually implement this endpoint, when we cover database
	// persistence later in the course.

	v := mux.Vars(r)

	if v["username"] == "EngineerKamesh" {

		mockPosts := make([]socialmedia.Post, 3)

		post1 := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post2 := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["happy"], "Tour of Go", "Check out the Tour of Go!", "https://tour.golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post3 := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["hopeful"], "Go Playground", "Check out the Go Playground!", "https://playground.golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})

		mockPosts = append(mockPosts, *post1)
		mockPosts = append(mockPosts, *post2)
		mockPosts = append(mockPosts, *post3)

		// 解析
		json.NewEncoder(w).Encode(mockPosts)

	} else {
		json.NewEncoder(w).Encode(nil)

	}

}

```

# model

```
// Package SocialMedia implements common functionality needed for social media web applications.
package socialmedia

import (
	"time"
)

//go:generate stringer -type=MoodState
type MoodState int

// All possible mood states.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

// AuditableContent types are meant to be embeded into types we want to keep a
// check on for auditing purposes
type AuditableContent struct {
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
}

// Post represents a Social Media Post type.
type Post struct {
	AuditableContent           // Embedded type
	Caption          string    `json:"caption"`
	MessageBody      string    `json:"messageBody"`
	URL              string    `json:"url"`
	ImageURI         string    `json:"imageURI"`
	ThumbnailURI     string    `json:"thumbnailURI"`
	Keywords         []string  `json:"keywords"`
	Likers           []string  `json:"likers"`
	AuthorMood       MoodState `json:"authorMood"`
}

// Map that holds the various mood states with keys to serve as
// aliases to their respective mood states.
var Moods map[string]MoodState

// The init() function is responsible for initializing the mood state
func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine}
}

// NewPost is responsible for creating an instance of the Post type.
func NewPost(username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post {

	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url, ImageURI: imageURI, ThumbnailURI: thumbnailURI, AuthorMood: mood, Keywords: keywords, AuditableContent: auditableContent}
}

```
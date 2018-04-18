package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, " dog 狗")
	case "/cat":
		io.WriteString(w, "猫猫猫....")

	}

}
func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}

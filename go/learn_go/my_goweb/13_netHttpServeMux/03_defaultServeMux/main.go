package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d)
	http.Handle("/cat", c) // 无效的路径必须有斜杠

	http.ListenAndServe(":8080", nil)
}

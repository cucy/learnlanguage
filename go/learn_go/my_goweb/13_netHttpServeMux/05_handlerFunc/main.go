package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

// HandlerFunc converts the handle func to a HandlerFunc which in turn
// implements the Handler interface. http.Handle takes a path and a handler

// handlerfunc转换处理功能的一handlerfunc反过来
// 实现处理器接口。句柄需要一条路径和一个处理程序。(需要一个请求的路径, handler)

func main() {
	http.Handle("/dog/", http.HandlerFunc(d)) // implements the Handler interface. http.Handle takes a path and a handler
	http.Handle("/cat", http.HandlerFunc(c))  // 无效路径

	http.ListenAndServe(":8080", nil)
}

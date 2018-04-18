package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results := dbAccess(ctx)
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
	fmt.Println("从数据库访问，context is From dbAccess, context is:", ctx)
	uid := ctx.Value("userID").(int) // 777       // assertion taking place here 断言发生在这里
	user_fname := ctx.Value("fname").(string)
	fmt.Println("fname: ", user_fname) // Bond
	return uid
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Println(w, ctx)
}

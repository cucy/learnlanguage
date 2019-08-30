package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	const addr = "localhost:8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
	})

	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			panic(err)
		}
	}()

	req, _ := http.NewRequest(http.MethodGet, "http://"+addr, nil)

	ctx, canc := context.WithTimeout(context.Background(), time.Second*2)
	defer canc()

	time.Sleep(time.Second)
	if _, err := http.DefaultClient.Do(req.WithContext(ctx)); err != nil {
		log.Fatal("超时", err)
	}
}

// 2019/08/28 09:28:28 超时Get http://localhost:8080: context deadline exceeded

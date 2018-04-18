package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	get()
}

func get() {
	res, err := http.Get("http://www.gutenberg.org/ebooks/56959")
	if err != nil {
		log.Fatal(err)

	}
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)

}

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func ex1() {
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."

	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	fmt.Println()
	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)
	fmt.Println()

	res, _ := http.Get("https://www.baidu.com/")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()

}

func main() {
	ex1()
}

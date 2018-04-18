package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func iocopy_ex_error_check() {
	msg := "这是一段话....."
	rdr := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()

	rdr2 := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, rdr2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()

	res, err := http.Get("http://www.mcleods.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, res.Body)
	if err := res.Body.Close(); err != nil {
		fmt.Println(err)
	}

}

func main() {
	iocopy_ex_error_check()
}

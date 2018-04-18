package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method) // POST
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // 抛出内部错误
			return
		}
		defer f.Close()

		fmt.Printf("\nfile:", f, "\nheader:", h, "\nerr", err)

		bs, err := ioutil.ReadAll(f) // 文件内容, 字符串数组型
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s) // 把提交的表单再返回到客户端
}

/*
new.txt 内容
测试文件,text
content.

*/

/*
output:

POST

file: {0xc0420f63f0}
header: &{new.txt map[Content-Disposition:[form-data; name="q"; filename="new.txt"] Content-Type:[text/plain]] 28 [230 181 139 232 175 149 230 150 135 228 187 182 44 116 101 120 116 32 13 10 99 111 110 116 101 110 116 46] }
err <nil>


*/

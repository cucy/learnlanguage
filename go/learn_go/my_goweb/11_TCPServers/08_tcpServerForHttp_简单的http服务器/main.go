package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write request
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	fmt.Println("==== 打印request信息 ===")
	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln) // 打印request信息

		if i == 0 {
			// first line in request with method
			firstLineFields := strings.Fields(ln)
			m := firstLineFields[0]
			url := firstLineFields[1]
			fmt.Println("***METHOD", m, "****URL", url)
		}
		if ln == "" {
			// headers are done, this is blank line before payload
			break
		}
		i++
	}
	fmt.Println("==== 打印request信息End ===")
}

func response(conn net.Conn) {
	body := `<DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Response</title>
	<body>
	<h1>Hello World!</h1>
	</body>
	</head>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

/*
==== 打印request信息 ===
GET / HTTP/1.1
***METHOD GET ****URL /
Host: 127.0.0.1:8080
Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.167 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9
Cookie: csrftoken=OKynYZhHEhhDYc8xE4qJizQV0OsZB0eS2W4pKomKFz4xmsa5aIvzOIWZhoWcKCIW

==== 打印request信息End ===
*/

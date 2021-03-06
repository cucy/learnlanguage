package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}

	li.Close()

}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	var i int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // 读取到的内容为空时 跳出循环
		}

		if i == 0 {
			fs := strings.Fields(line)
			method := fs[0]
			url := fs[1]
			fmt.Println("****METHOD:", method)
			fmt.Println("****URL:", url)
		}
		fmt.Println(line)
		i++
	}

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

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	fmt.Fprint(conn, body)
	// io.WriteString(conn, "I see you connected!")

	fmt.Println("Code got here")

	conn.Close()
}

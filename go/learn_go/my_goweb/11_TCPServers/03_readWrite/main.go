package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		fmt.Fprintf(conn, "I heard you say: %s\n ", line)
	}
	conn.Close()
}

/*
[c:\~]$ telnet 127.0.0.1 8080


Connecting to 127.0.0.1:8080...
Connection established.
To escape to local shell, press Ctrl+Alt+].
a
I heard you say: a
                   s
I heard you say: s
                   dd


*/

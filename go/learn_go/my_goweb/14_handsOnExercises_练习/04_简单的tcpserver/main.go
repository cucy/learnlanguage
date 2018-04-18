package main

import (
	"fmt"
	"log"
	"net"
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
	fmt.Fprintln(conn, "I can hear you!")
	conn.Close()
}

/*

[c:\~]$ telnet 127.0.0.1 8080

*/

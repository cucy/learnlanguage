// In addition to this program, also need to run 2_read that reads from a connection to the server
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(conn, "我已经接收到信息!!")
	conn.Close()

}

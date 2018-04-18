// In addition to running this file, you also need to run main.go in 01_write
// in order to have a tcp server writing output which this program can read
// 除了运行这个文件，你还需要在运行main.go 01_write
// 为了让TCP服务器编写输出，这个程序可以读取
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("bs:", bs)
	fmt.Println(string(bs))
}

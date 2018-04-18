package main

import "fmt"

func no_defer_hello() {
	fmt.Print("hello ")
}
func no_defer_world() {
	fmt.Println("world")
}
func main() {
	no_defer_hello()
	no_defer_world()
	/*hello world*/
	//	 没有defer, 程序从上往下执行

	//	with  defer
	defer world() // 最后执行
	hello()
	/*hello world*/
}

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

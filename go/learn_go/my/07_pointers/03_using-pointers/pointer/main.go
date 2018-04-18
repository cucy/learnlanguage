package main

import "fmt"

func zero(z *int) {
	fmt.Println(z)             // 0xc042052080  zero中因为z的类型为*int int型指针 值就是X的地址
	fmt.Println("指针z自己地址", &z) // 指针z自己地址 0xc042072020
	*z = 0                     // 改值
}
func main() {
	X := 5
	fmt.Println(&X) // main函数中 X变量的地址为 0xc042052080

	zero(&X)       // 传地址
	fmt.Println(X) // 0
}

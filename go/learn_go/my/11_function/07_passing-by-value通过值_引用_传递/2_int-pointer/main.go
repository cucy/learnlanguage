package main

import "fmt"

func main() {
	age := 666
	fmt.Println("age address", &age) // age address 0xc042052080

	change_me(&age)

	fmt.Println(&age) // 0xc042052080
	fmt.Println(age)  //  888
}

func change_me(z *int) {
	fmt.Println("z address ", z) // z address  0xc042052080
	fmt.Println(*z)              // 获取0xc042052080所存的值

	*z = 888 // 改变0xc042052080所存的值

	fmt.Println("z address", z) // 取地址    z address 0xc042052080
	fmt.Println("z value", *z)  // z value 888

}

/*
当需要在不同的函数修改同一个值时,可以使用指针

指针指向的是值所在的地址,当修改该地址所指向的值,该值会全局修改,而不是得到一个副本(病句)


*/

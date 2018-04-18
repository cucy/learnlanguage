package main

import "fmt"

func change_me(z int) {
	fmt.Println(z) // 888
	z = 666
}
func main() {
	age := 888
	change_me(age)
	fmt.Println(age) // 888
}

// when changeMe is called on line 8
// the value 44 is being passed as an argument
// 当在第11行调用change_me时,
// 值44作为参数传递。

/*因为是传值所以不会修改原来的值age*/

package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42

}

/*
	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value


	// 这是有用的
	// 我们可以传递一个内存地址，而不是一堆值(我们可以通过引用)
	// 然后我们仍然可以改变存储在那个内存地址中的值。
	// 这使我们的programs更有表现力。
	// 我们不需要传递大量的数据。
	// 我们只需要传递地址。

	// 顺便说一下，一切都是通过值传递的。
	// 当我们传递一个内存地址时，我们正在传递一个值。

*/

// 引用
package main

import "fmt"

func main() {

	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)

	de_referencing()
}

/*

43
0xc0420080a8
0xc0420080a8

*/

/*

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int

// 上面的代码使b成为一个指向内存地址的指针。
//  b类型为“int指针”
// * *是类型的一部分，b是类型*int类型。
*/

func de_referencing() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case

	// b是一个int指针;
	// b指向存储一个int的内存地址。
	// 要查看内存地址的值，在b前面加上一个*。
	// 这就是所谓的“引用”。
	// 在这种情况下*是一个操作符。
}

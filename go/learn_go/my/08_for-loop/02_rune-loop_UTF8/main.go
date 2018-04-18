package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := "a"

	fmt.Println(foo)
	fmt.Printf("%T \n", foo)

	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}

/*
NOTE:
Some operating systems (Windows) might not print characters where i < 256

If you have this issue, you can use this code:

fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))

UTF-8 is the text coding scheme used by Go.

UTF-8 works with 1 - 4 bytes.

A byte is 8 bits.

[]byte deals with bytes, that is, only 1 byte (8 bits) at a time.

[]int32 allows us to store the value of 4 bytes, that is, 4 bytes * 8 bits per byte = 32 bits.

注意:
有些操作系统(Windows)可能不会输出i < 256的字符。

如果您有这个问题，您可以使用以下代码:

fmt。Println(“-”,字符串(我),“-”,[]int32(字符串(i)))

UTF-8是Go所使用的文本编码方案。

UTF-8可以使用1 - 4字节。

一个字节是8位。

字节处理字节，也就是说，一次只处理一个字节(8位)。

[]int32允许我们存储4字节的值，即4字节* 8位/字节= 32位。
*/

package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "我很怀念刚认识那会,大家都有点谨慎和真诚."
		fmt.Println(y)
	} // 大括号亦是一个块级别的作用域
	//fmt.Println(y) //  undefined: y  编译错误
}

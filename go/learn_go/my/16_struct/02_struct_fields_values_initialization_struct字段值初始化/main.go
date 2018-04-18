package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"张", "三", 99}
	p2 := person{"李", "四", 99}
	fmt.Println(p1.first, p1.last, p1.age) // 张 三 99
	fmt.Println(p2.first, p2.last, p2.age) //  李 四 99
}

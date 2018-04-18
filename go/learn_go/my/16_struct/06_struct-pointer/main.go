package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"张三", 100}
	fmt.Println(p1)        // &{张三 100}
	fmt.Printf("%T\n", p1) // *main.person
	fmt.Println(p1.name)   // 张三
	fmt.Println(p1.age)    // 100
}

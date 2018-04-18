package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type duble_zero struct {
	person
	First   string
	License bool
}

func ex1() {
	p1 := duble_zero{
		person: person{
			First: "张",
			Last:  "三",
			Age:   111,
		},
		First:   "duble_zero first field",
		License: false,
	}
	p2 := duble_zero{
		person: person{
			First: "王",
			Last:  "西瓜",
			Age:   111123123,
		},
		First:   "王师傅",
		License: true,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	//	内部类型的字段和方法被提升为outer类型。
	fmt.Println(p1.First, p1.person.First) // duble_zero first field 张
	fmt.Println(p2.First, p2.person.First) // 王师傅 王
}

func main() {

	//ex1()
	ex2()
}

func (p person) Greeting() {
	fmt.Println("我只是一个普通人.")
}
func (dz duble_zero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func ex2() {
	p1 := person{
		First: "张",
		Last:  "无",
		Age:   100,
	}

	p2 := duble_zero{
		person: person{
			First: "James Bond",
			Last:  "无",
			Age:   23,
		},
		License: true,
	}

	p1.Greeting() // 我只是一个普通人.
	p2.Greeting()
	p2.person.Greeting()

}

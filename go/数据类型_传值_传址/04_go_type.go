package main

import "fmt"

// 值类型
func modify_str(s string) string {
	s = s + s
	return s
}

// 引用类型
func reference_type(m map[string]int) int {
	m["张三"] = 1000
	return 0
}

//  结构类型
type person struct {
	age  int
	name string
}

func modify_struct(p *person) {
	p.age = p.age + 10
}

// 自定义类型
type Duration int64

var i Duration = 100
var j int64 = 100

func main() {

	// 基本类型因为是拷贝的值
	name := "张三"
	fmt.Println(modify_str(name))
	fmt.Println(name)

	// 引用类型
	ages := map[string]int{"张三": 20}
	fmt.Println(ages)
	reference_type(ages)
	fmt.Println(ages)

	//	 结构类型
	jim := person{10, "Jim"} // 如果不指定关键字必须按位置传入
	fmt.Println(jim)
	jim1 := person{name: "Jim", age: 10} // 指定关键字后位置可以随意
	fmt.Println(jim1)

	jim3 := person{10, "Jim3"}
	fmt.Println("结构类型jim3修改前的值: ", jim3)
	modify_struct(&jim3)
	fmt.Println("结构类型jim3修改后的值:", jim3)
}


# for循环引用问题


```go
package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func parse_stu() {
	m := make(map[string]*Student)

	stu := []Student{
		{Name: "su", Age: 100},
		{Name: "li", Age: 12},
		{Name: "mo", Age: 16},
	}

	for _, s := range stu {
		fmt.Printf("s的地址: %p\n", &s)
		m[s.Name] = &s // 有bug

		/*  重新赋值才会申请新的内存 */
		//   s1:=s
		// m[s1.Name] = &s1
	}
	fmt.Println()
	fmt.Printf("%#v\n\n", m)
	for k, v := range m {
		fmt.Printf("k,v的地址: %p  %p\n", &k, &v)
		fmt.Println(k, v)
	}

}
func main() {
	parse_stu()
}



/*

出现此原因是在语句块, 只第一次申请内存, 而后再次循环只会赋新值, 而不会申请新的地址, 所有使用引用的时候要注意

s的地址: 0xc000050420
s的地址: 0xc000050420
s的地址: 0xc000050420

map[string]*main.Student{"li":(*main.Student)(0xc000050420), "mo":(*main.Student)(0xc000050420), "su":(*main.Student)(0xc000050420)}

k,v的地址: 0xc0000461f0  0xc000086020
li &{mo 16}
k,v的地址: 0xc0000461f0  0xc000086020
mo &{mo 16}
k,v的地址: 0xc0000461f0  0xc000086020
su &{mo 16}


*/



```

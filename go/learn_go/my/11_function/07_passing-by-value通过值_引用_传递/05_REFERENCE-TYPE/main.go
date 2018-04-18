package main

import "fmt"

func main() {
	m := make(map[string]int)
	change_me(m)
	fmt.Println(m["Age"]) // 100
}

func change_me(z map[string]int) {
	z["Age"] = 100
}

/*
Allocation with make
Back to allocation. The built-in function make(T, args)
serves a purpose different from new(T). It creates slices, maps, and channels only,
and it returns an initialized (not zeroed) value of type T (not *T). The reason for
the distinction is that these three types represent, under the covers, references to data structures
that must be initialized before use. A slice, for example, is a three-item descriptor containing
a pointer to the data (inside an array), the length, and the capacity, and until those items are initialized,
the slice is nil. For slices, maps, and channels, make initializes the internal data structure and prepares
the value for use.


分配与使
重新分配。内置函数make(T, args)
服务于不同于新(T)的目的。它只创建切片、映射和通道，
它返回一个初始化(不是零)类型T (not *T)的值。的原因
区别在于，这三种类型分别代表了数据结构的引用。
必须在使用之前对其进行初始化。例如，slice是包含三个项目的描述符。
指向数据(数组内)、长度和容量的指针，直到这些项被初始化，
片是零。对于切片、映射和通道，可以初始化内部数据结构并准备。
使用的值。
*/

// https://golang.org/doc/effective_go.html#allocation_make

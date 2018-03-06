package main

import "fmt"

func main() {

	// 一个是指向底层数组的指针，
	// 一个是切片的长度，
	// 一个是切片的容量
	slice := make([]int, 5) // 指定切片的容量为 5, 长度为5
	fmt.Printf("%v \n", slice)

	slice2 := make([]int, 5, 10)
	fmt.Printf(
		// slice2容量为:10, slice2长度:5
		"slice2容量为:%v, slice2长度:%v \n",
		cap(slice2), len(slice2))

	// 第二种方式
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v \n", slice3)

	// 初始化第五个值为1
	slice_init5 := []int{4: 1}
	fmt.Printf("%v \n", slice_init5) // [0 0 0 0 1]

	// 注意数组和切片的差异
	/*
		数组
		array:=[5]int{4:1}
		切片
		slice:=[]int{4:1}
	*/

	/*
			// nil切片
			var nilSlice []int
			// 空切片
			slice := []int{}

		// nil切片表示不存在的切片，而空切片表示一个空集合，它们各有用处
	*/

	// 访问
	fmt.Print("\nslice访问 \n")
	slice_ex := []int{1, 2, 3, 4, 5}
	slice_ex1 := slice_ex[:]
	slice_ex2 := slice_ex[0:]
	slice_ex3 := slice_ex[:5]
	slice_ex4 := slice_ex[1:3]   // 左开右闭原则
	slice_ex5 := slice_ex[0:2:4] // slice_ex5 长度为 2-0=2  容量为 4-0 = 2

	fmt.Println(slice_ex1)
	fmt.Println(slice_ex2)
	fmt.Println(slice_ex3)
	fmt.Println(slice_ex4)

	fmt.Println("slice_ex5:", slice_ex5, "长度为:", len(slice_ex5), "容量为:", cap(slice_ex5))
	slice_ex5 = append(slice_ex5, 10)
	fmt.Println("slice_ex5:", slice_ex5, "长度为:", len(slice_ex5), "容量为:", cap(slice_ex5)) // slice_ex5: [1 2 10] 长度为: 3 容量为: 4
	//fmt.Println(slice_ex5[4:])
	//fmt.Println(slice_ex5[5:])  // 报错 slice bounds out of range
	slice_ex5 = append(slice_ex5, 10, 11, 12, 13)
	fmt.Println("slice_ex5:", slice_ex5, "长度为:", len(slice_ex5), "容量为:", cap(slice_ex5)) // slice_ex5: [1 2 10 10 11 12 13] 长度为: 7 容量为: 8

	slice_ex5 = append(slice_ex5, slice_ex3...)
	fmt.Println("slice_ex5:", slice_ex5, "长度为:", len(slice_ex5), "容量为:", cap(slice_ex5)) // slice_ex5: [1 2 10 10 11 12 13 1 2 10 4 5] 长度为: 12 容量为: 16

	// 迭代
	slice_for := []int{1, 2, 3, 11}
	for k, v := range slice_for {
		fmt.Printf("key:%v, value:%v\n", k, v)
	}

	//  传统for循环取值
	slice_for_old := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("值:%d\n", slice_for_old[i])
	}

	// 在函数间传递切片
	slice_for_fun := []int{1, 2, 3, 4, 5}
	fmt.Printf("\n")
	fmt.Printf("%p\n", &slice_for_fun)
	modify(slice_for_fun)
	fmt.Println(slice_for_fun)
	// http://www.flysnow.org/2017/03/14/go-in-action-go-slice.html
}

func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}

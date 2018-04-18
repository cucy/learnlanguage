package main

import "fmt"

func init_condition_post01() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func nested02() {

	//	嵌套
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}

}

func while_loop_in_go() {
	// go语言中没有while关键字
	i := 0
	for i < 10 {
		fmt.Println("while_loop_in_go", i)
		i++
	}
}

func while_true() {
	//	 死循环
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 10 {
			break
		}
	}

}

func for_break() {
	// break
	i := 0
	for {
		fmt.Println("for_break", i)
		if i >= 10 {
			break
		}
		i++
	}
}

func for_continue06() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println("for_continue06", i)
		if i >= 50 {
			break
		}
	}
}

func main() {
	init_condition_post01()
	nested02()
	while_loop_in_go()
	while_true()
	for_break()
	for_continue06()
}

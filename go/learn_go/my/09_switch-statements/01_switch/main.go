package main

import "fmt"

func switch_default() {
	switch "Mhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")

	}
}

func switch_fallthrough() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough // 不管是不是true 一定会执行
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough // 不管是不是true 一定会执行
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}

func switch_multiple_evals() {
	// 多个条件
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}

// 没有表达式 04_no-expression
func switch_no_expression() {
	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")

	}
	/*
				  expression not needed
				  -- if no expression provided, go checks for the first case that evals to true
				  -- makes the switch operate like if/if else/else
				  cases can be expressions

		    不需要表达
			如果没有提供表达式，则检查第一个事件是否为真。
			——使开关的运行方式类似于if/if else/else。
			情况下可以表达
	*/

}
func main() {

	switch_default()
	fmt.Println()
	switch_fallthrough()
	fmt.Println()
	switch_multiple_evals()

	fmt.Println("switch_no_expression")
	switch_no_expression()
}

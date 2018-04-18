package main

import "fmt"

// 全局变量
var Myname = "Centos"

/* ---------------------- 定义变量 ---------------------- */
//
/* ---------------------- 定义变量 ---------------------- */
func declare_variable() {
	/* 先定义 后初始化 */
	var message string
	message = "hello world"
	fmt.Println(message)
	fmt.Println(Myname)
	fmt.Println()
}

/* ---------------------- 一次定义多个变量 ---------------------- */
//
/* ---------------------- 一次定义多个变量 ---------------------- */
func declare_some_variable() {
	var message string
	var a, b, c int // 一次定义多个变量
	a = 1           // 赋值
	message = "hello"
	fmt.Println(message, a, b, c) // hello 1 0 0
	fmt.Println()
}

/* ---------------------- 定义且赋值 自动推导 ---------------------- */
//
/* ---------------------- 定义且赋值 自动推导---------------------- */
func init_variable_at_once() {
	var message = "hello world!"
	var a, b, c int = 1, 2, 3

	var a1, b1, c1 = 11, 22, 33 // infer type 自动推导不指定变量类型

	var mix1, mix2, mix3 = 1, false, 3 // 自动推导也可以是不同种类型数据

	fmt.Println(message, a, b, c) // hello world! 1 2 3
	fmt.Println()
	fmt.Println(a1, b1, c1)       // 11 22 33
	fmt.Println(mix1, mix2, mix3) // 1 false 3
	fmt.Println()

}

/* ---------------------- 简短声明 ---------------------- */
//
/* ---------------------- 简短声明---------------------- */
func variable_init_shorthand() {
	// 1.简短声明只能在函数内部
	// 2.变量名没有被声明过

	message := "hello world!"
	a, b, c := 1, false, 300
	d := 4
	e := true
	fmt.Println(message, a, b, c, d, e) // hello world! 1 false 300 4 true
	fmt.Println()

}

/* ---------------------- all together 汇总---------------------- */
//
/* ---------------------- all together 汇总---------------------- */

// package scope 包级别作用域
var aaa = "这个是存储在变量aaa中"

// package scope 包级别作用域
var bbb, ccc string = "保存在bbb", "保存在ccc"

// package scope 包级别作用域
var d string

func all_together() {
	d = "stored in d" // declaration above; assignment here; package scope
	var e = 42        // function scope - subsequent variables have func scope:
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm' // 单引号
	n := "n"                             // double quotes
	o := `o`                             // 反勾号 like  double quotes

	fmt.Println("\n==== 汇总 ===")
	fmt.Println("aaa - ", aaa)
	fmt.Println("bbb - ", bbb)
	fmt.Println("ccc - ", ccc)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
}

func main() {
	declare_variable()

	declare_some_variable()

	init_variable_at_once()

	variable_init_shorthand()

	all_together()
}

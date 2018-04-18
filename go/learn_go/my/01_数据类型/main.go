package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 整形 decimal
	fmt.Println(42)

	// 二进制 binary
	fmt.Printf("%d ==> %b \n", 42, 42)

	// 十六进制 hexadecimal
	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)

	for i := 0; i < 127; i++ {
		// ASCII
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}

	fmt.Println("===================================")
	a := "this is stored in the variable a"
	b := 42
	c, d, e, f := 44.7, true, false, 'm' // single quotes
	g := "g"                             // double quotes
	h := `h`                             // back ticks

	fmt.Println("a - ", reflect.TypeOf(a), " - ", a)
	fmt.Println("b - ", reflect.TypeOf(b), " - ", b)
	fmt.Println("c - ", reflect.TypeOf(c), " - ", c)
	fmt.Println("d - ", reflect.TypeOf(d), " - ", d)
	fmt.Println("e - ", reflect.TypeOf(e), " - ", e)
	fmt.Println("f - ", reflect.TypeOf(f), " - ", f)
	fmt.Println("g - ", reflect.TypeOf(g), " - ", g)
	fmt.Println("h - ", reflect.TypeOf(h), " - ", h)
	fmt.Printf("h -  %T\n", h)
}

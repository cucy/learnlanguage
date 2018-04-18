package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	aa = iota // 0
	bb        // 1
	cc        // 2
)

const (
	a1 = iota // 0
	b1        // 1
	c1        // 2

	//	一个作用域,不影响到其他
)

const (
	d1 = iota // 0
	e1        // 1
	f1        // 2
	//	一个作用域,不影响到其他
)

const (
	_  = iota      // 0
	z1 = iota * 10 // 1 * 10
	z2 = iota * 10 // 2 * 10

)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (2*10)
	GB = 1 << (iota * 10) // 1 << (3*10)
	TB = 1 << (iota * 10) // 1 << (4*10)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println()
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	fmt.Println()
	fmt.Println(z1)
	fmt.Println(z2)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("KB %b\t", KB)
	fmt.Printf("KB %d\n", KB)
	fmt.Printf("MB %b\t", MB)
	fmt.Printf("MB %d\n", MB)
	fmt.Printf("GB %b\t", GB)
	fmt.Printf("GB %d\n", GB)
	fmt.Printf("TB %b\t", TB)
	fmt.Printf("TB %d\n", TB)
}

/*
0
1
2

0
1
2

10
20
binary		decimal
KB 10000000000	KB 1024
MB 100000000000000000000	MB 1048576
GB 1000000000000000000000000000000	GB 1073741824
TB 10000000000000000000000000000000000000000	TB 1099511627776
*/

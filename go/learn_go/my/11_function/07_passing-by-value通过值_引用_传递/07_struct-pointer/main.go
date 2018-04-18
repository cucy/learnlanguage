package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	u1 := user{"centos", 1990}
	fmt.Println("&u1 ", &u1) // &u1  &{centos 1990}

	change_me(&u1)

	fmt.Println(u1)       // {red hat 1990}
	fmt.Println(&u1.name) // 0xc04204c3e0

}

func change_me(z *user) {
	fmt.Println(z)       //  &{centos 1990}
	fmt.Println(&z.name) // 0xc04204c3e0
	z.name = "red hat"

	fmt.Println("after z ", z) // after z  &{red hat 1990}
	fmt.Println(&z.name)       // 0xc04204c3e0
}

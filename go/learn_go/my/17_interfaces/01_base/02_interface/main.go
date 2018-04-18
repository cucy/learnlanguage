package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)        // {10}
	fmt.Println(z.area()) // 100
}

func main() {
	s := square{10}
	fmt.Printf("%T\n", s) // main.square
	info(s)
}

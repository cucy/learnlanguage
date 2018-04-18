package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

// another shape 另一个形状
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

// which implements the shape interface 哪个实现了形状接口?
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}

/*
{10}
100
{5}
78.53981633974483

*/

package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}

/*
0  -  {{0 0 } 0 0}
1  -  {{0 0 } 0 0}
2  -  {{0 0 } 0 0}
3  -  {{0 0 } false}
4  -  {{0 0 } false}
5  -  {{0 0 } false}
6  -  {{0 0 } 0}
7  -  {{0 0 } 0}
8  -  {{0 0 } 0}

*/

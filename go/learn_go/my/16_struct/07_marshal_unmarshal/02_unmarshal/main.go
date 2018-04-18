package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func ex1() {
	var p1 person
	fmt.Println(p1.First) //
	fmt.Println(p1.Last)  //
	fmt.Println(p1.Age)   // 0

	bs := []byte(`{"First":"张", "Last":"三", "Age":20}`)

	json.Unmarshal(bs, &p1)
	fmt.Println("--------------")
	fmt.Println(p1.First)   // 张
	fmt.Println(p1.Last)    // 三
	fmt.Println(p1.Age)     // 20
	fmt.Printf("%T \n", p1) // main.person

}

func main() {
	//ex1()
	ex2_tags()
}

type person1 struct {
	First string
	Last  string
	Age   int `json:"年龄"`
}

func ex2_tags() {
	var p1 person1

	fmt.Println(p1.First) //
	fmt.Println(p1.Last)  //
	fmt.Println(p1.Age)   // 0

	bs := []byte(`{"First":"张", "Last":"三", "年龄":20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1.First)   // 张
	fmt.Println(p1.Last)    // 三
	fmt.Println(p1.Age)     // 20
	fmt.Printf("%T \n", p1) // main.person1
}

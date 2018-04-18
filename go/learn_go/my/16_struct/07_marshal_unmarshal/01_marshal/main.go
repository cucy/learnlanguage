package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int // 不导出
}

func encode_ex1() {
	p1 := person{"张", "三", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)         // [123 34 70 105 114 115 116 34 58 34 229 188 160 34 44 34 76 97 115 116 34 58 34 228 184 137 34 44 34 65 103 101 34 58 50 48 125]
	fmt.Printf("%T \n", bs) // []uint8
	fmt.Println(string(bs)) // {"First":"张","Last":"三","Age":20} // notExported并没有导出

}

func main() {
	//encode_ex1()
	//ex2()
	tags()
}

type person_un_exported struct {
	first string
	last  string
	age   int
}

func ex2() {
	p1 := person_un_exported{"张", "三", 200}
	fmt.Println(p1) // {张 三 200}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs)) // {} 小写不会导出
}

type person_tag struct {
	First string
	Last  string `json: "-"`
	Age   int    `json: "wisdom score"`
}

func tags() {
	p1 := person_tag{"张", "三", 200}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs)) // {"First":"张","Last":"三","Age":200}
}

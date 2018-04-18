package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func ex_encode() {
	p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1) // {"First":"James","Last":"Bond","Age":20}
}

func ex_decode() {
	var p1 person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First) // James
	fmt.Println(p1.Last)  // Bond
	fmt.Println(p1.Age)   // 20

}
func main() {
	ex_decode()
}

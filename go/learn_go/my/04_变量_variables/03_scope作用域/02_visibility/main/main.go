package main

import (
	"fmt"
	"learn_go/my/04_变量_variables/03_scope作用域/02_visibility/vis"
)

func main() {

	vis.Print_var()
	fmt.Println("外部包引用-->", vis.My_name)

}

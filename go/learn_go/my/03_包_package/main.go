package main

import (
	"fmt"
	testvar "learn_go/my/03_包_package/01_变量可见与不可见"
	stringutil "learn_go/my/03_包_package/02_stringutil"
)

func main() {
	fmt.Println(testvar.Phone) // 引用外部包变量
	// fmt.Println(testvar.my_phone)  //  引用外部包变量 报错, 无法引用

	fmt.Println(stringutil.My_name)        // 引用外部包变量
	fmt.Println(stringutil.Reverse("ABC")) // 引用外部包函数

}

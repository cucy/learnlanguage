package main

// http://www.flysnow.org/2017/03/23/go-in-action-go-map.html
import (
	"fmt"

	"reflect"
	"sort"
)

func main() {
	user := make(map[string]int)
	user["age"] = 100
	fmt.Print(user)

	dict := map[string]int{"张三": 43}
	fmt.Print("\n", dict)

	dict1 := map[string]int{"张三": 43, "李四": 50}
	fmt.Print("\n同时初始化多个值: \n", dict1)

	// 空map
	null_map := map[string]int{}
	fmt.Print("\n", null_map)

	//var dict1 map[string]int
	fmt.Printf("\n \n")
	var dict_1 map[string]int
	dict_1 = make(map[string]int)
	dict_1["张三"] = 43
	fmt.Println(dict_1)

	//	===================

	dict11 := make(map[string]int)
	dict11["张三"] = 43 // 存在则修改

	// 获取值
	get_age := dict11["张三"]
	fmt.Printf("\n get_age:%v", get_age)

	// 判断key是否存在
	age2, exists := dict["李四"] // 不存在以下不会执行
	if exists {
		fmt.Printf("\nage2:%v", age2)
	}

	// 删除key
	delete(dict, "张三")

	// 遍历map
	fmt.Printf("\n \n")
	dict12 := map[string]int{"张三": 43}
	for key, value := range dict12 {
		fmt.Println(key, value)
	}

	//  值排序
	fmt.Printf("\n \n")
	dict123 := map[string]int{"王五": 60, "张三": 43, "李四": 1}
	var names []string //  存放map中的key
	fmt.Printf("%v \n", reflect.TypeOf(names))

	for key := range dict123 {
		// 把key放到切片中
		names = append(names, key)

	}
	sort.Strings(names) // 排序
	for v, key := range names {
		fmt.Println(v, key, dict123[key])
	}

	//	在函数间传递Map

	dict_for_func := map[string]int{"王五": 60, "张三": 43}
	fmt.Println("修改前的值为", dict_for_func["张三"])
	modify_map(dict_for_func)
	fmt.Println("修改后的值为", dict_for_func["张三"])

}
func modify_map(dict map[string]int) {
	dict["张三"] = 10
}

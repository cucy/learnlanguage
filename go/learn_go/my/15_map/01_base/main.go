package main

import "fmt"

// var 声明
func var_init_map_string() {
	var my_greeting map[string]string
	fmt.Println(my_greeting)        // map[]
	fmt.Println(my_greeting == nil) // true

}

// make  声明
func make_init_map() {
	var my_greeting = make(map[string]string)
	my_greeting["张三"] = "中国"
	my_greeting["金字塔"] = "埃及"
	fmt.Println(my_greeting) // map[张三:中国 金字塔:埃及]

}

// 自动推导定义 声明
func shortand_make() {
	my_greeting := make(map[string]string)
	my_greeting["张三"] = "中国"
	my_greeting["金字塔"] = "埃及"
	fmt.Println(my_greeting) // map[张三:中国 金字塔:埃及]
}

// 声明且初始化 空值
func shorthand_composite_literal() {
	my_greeting := map[string]string{}
	my_greeting["张三"] = "中国"
	my_greeting["金字塔"] = "埃及"
	fmt.Println(my_greeting) // map[张三:中国 金字塔:埃及]
}

// 声明且初始化且赋值
func shorthand_composite_literal1() {
	my_greeting := map[string]string{
		"张三":  "中国",
		"金字塔": "埃及",
	}
	fmt.Println(my_greeting["张三"]) // 中国
}

// add 向map添加元素
func add_ele_to_map() {
	my_greeting := map[string]string{
		"张三":  "中国",
		"金字塔": "埃及",
	}
	my_greeting["张国荣"] = "香港"
	fmt.Println(my_greeting) // map[张三:中国 金字塔:埃及 张国荣:香港]
}

// len 获取map长度
func get_map_length() {
	my_greeting := map[string]string{
		"张三":  "中国",
		"金字塔": "埃及",
	}
	my_greeting["张国荣"] = "香港"
	fmt.Println(len(my_greeting)) // 3
}

// update 修改map的值
func update_map_value() {
	my_greeting := map[string]string{
		"张三":  "中国",
		"金字塔": "埃及",
	}
	fmt.Println(my_greeting) // map[张三:中国 金字塔:埃及]
	my_greeting["金字塔"] = "不在中国"
	fmt.Println(my_greeting) // map[金字塔:不在中国 张三:中国]

}

// delete map element 删除map中的元素
func delete_element_map() {
	my_greeting := map[string]string{
		"张三":  "中国",
		"金字塔": "埃及",
		"11":  "22",
	}
	fmt.Println(my_greeting) // map[张三:中国 金字塔:埃及 11:22]
	delete(my_greeting, "11")
	fmt.Println(my_greeting) // map[金字塔:埃及 张三:中国]

}

// is exists 判断某元素是否在map中
func element_is_exist_in_map() {
	my_stu := map[int]string{
		0: "张三",
		1: "李四",
		2: "王五",
	}

	fmt.Println(my_stu, "\n")
	//delete(my_stu, 2)

	if val, exists := my_stu[2]; exists {
		fmt.Println("key存在于map中.")
		fmt.Println("val: ", val)       // val:  王五
		fmt.Println("exists: ", exists) // exists:  true
	} else {
		fmt.Println("key不存在于map中.")
		fmt.Println("val: ", val)       // val:
		fmt.Println("exists: ", exists) // exists:  false
	}

}

// delete 删除不存在的key不会报错
func deleting_entry_no_error() {
	my_stu := map[int]string{
		0: "张三",
		1: "李四",
		2: "王五",
	}
	fmt.Println(my_stu) // map[0:张三 1:李四 2:王五]
	delete(my_stu, 88)
	fmt.Println(my_stu) // map[0:张三 1:李四 2:王五]
}

// delete key存在则删除
func delete_entry_if_exists_of_map() {
	my_stu := map[int]string{
		0: "张三",
		1: "李四",
		2: "王五",
	}
	fmt.Println(my_stu, "\n")

	if val, exists := my_stu[0]; exists {
		delete(my_stu, 7)
		fmt.Println("val: ", val)       // val:  张三
		fmt.Println("exists: ", exists) // exists:  true
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)       // val:
		fmt.Println("exists: ", exists) // exists:  false
	}
}

// for 循环 map
func for_loop_map() {
	my_stu := map[int]string{
		0: "张三",
		1: "李四",
		2: "王五",
	}
	for k, v := range my_stu {
		fmt.Printf("key:%d, value:%s \n", k, v)
	}
	/*
	   key:0, value:张三
	   key:1, value:李四
	   key:2, value:王五
	*/
}

// init

func main() {
	for_loop_map()
}

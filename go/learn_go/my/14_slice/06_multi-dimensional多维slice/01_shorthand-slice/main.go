package main

import "fmt"

func shorthand_slice() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)        // []
	fmt.Println(students)       // []
	fmt.Println(student == nil) // false

}
func var_slice() {
	var student []string
	var students [][]string
	fmt.Println(student)        // []
	fmt.Println(students)       // []
	fmt.Println(student == nil) // true

}

func make_slice() {
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)        // [                                  ]
	fmt.Println(students)       // [[] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] []]
	fmt.Println(student == nil) // false
}

func make_slice_duowei() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Todd"
	// student = append(student, "Todd")
	fmt.Println(student)  // [Todd                                  ]
	fmt.Println(students) // [[] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] []]
}

func make_slice_of_string() {
	var records [][]string

	// student 1
	stu1 := make([]string, 4)
	stu1[0] = "Foster"
	stu1[1] = "Nathan"
	stu1[2] = "100.00"
	stu1[3] = "74.00"
	// store the record
	records = append(records, stu1)

	// student 2
	stu2 := make([]string, 4)
	stu2[0] = "张三"
	stu2[1] = "中国"
	stu2[2] = "10.00"
	stu2[3] = "4.00"
	// store the record
	records = append(records, stu2)

	fmt.Println(records) // [[Foster Nathan 100.00 74.00] [张三 中国 10.00 4.00]]

}

func slice_of_slice_of_int() {
	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions) // [[0 1 2 3] [0 1 2 3] [0 1 2 3]]
}

func int_slice_plus_plus() {
	my_slice := make([]int, 1)
	fmt.Println(my_slice[0])

	my_slice[0] = 7
	fmt.Println(my_slice[0]) // 7
	my_slice[0]++
	fmt.Println(my_slice[0]) // 8

}

func main() {
	int_slice_plus_plus()
}

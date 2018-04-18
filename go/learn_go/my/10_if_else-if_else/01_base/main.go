package main

import "fmt"

func base_if() {
	if true {
		fmt.Println("this true")
	}
	if false {
		fmt.Printf("this did not run ")
	}
}

func if_not_exclamation() {
	//	 02_not-exclamation ! 感叹号

	if !true {
		fmt.Println("this did not run")
	}
	if !false {
		fmt.Printf("this  run ")
	}

}

func if__init_statement() {
	//	带有初始化的条件循环
	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}

}

func if_init_statement_error_invalid_code() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// fmt.Println(food)
	// 错误的引用, 因为food的作用域是在if块里
}

func if_else() {
	if false {
		fmt.Println("first print statement")
	} else {
		fmt.Println("second print statement")
	}
}

func if_elseif_else() {
	if false {
		fmt.Println("first print statement")
	} else if true {
		fmt.Println("second print statement")
	} else {
		fmt.Println("third print statement")
	}
}

func if_elseif_elseif_esle() {
	if false {
		fmt.Println("first print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("ahahaha print statement")
	} else {
		fmt.Println("third print statement")
	}
}

func divisibleByThree() {
	// 可以被3整除的数
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// 1000 内 3/5倍数
func three_five() {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}

func main() {
	base_if()

	fmt.Println("\nif_not_exclamation")
	if_not_exclamation()

	fmt.Printf("\nif__init_statement\n")
	if__init_statement() // Chocolate

	fmt.Println("if_else")
	if_else()

	fmt.Printf("\n")
	fmt.Printf("if_elseif_else\n")
	if_elseif_else()

	fmt.Printf("\nif_elseif_elseif_esle\n")
	if_elseif_elseif_esle()

	fmt.Printf("\ndivisibleByThree可以被3整的数\n")
	divisibleByThree()
}

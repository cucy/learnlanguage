# for

```go
	for { // 无限循环
		block
	}

	for boolExpression { // while
		block
	}

	for optionsPreStatement; booleanExpress; optionalPOstStatement { //
		block
	}

	for index, char := range aString { // 一个字符一个字符迭代字符串
		block
	}

	for index := range aString { // 一个字符一个字符迭代字符串
		block //  char,size := utf8.DecodeRuneInString(aString[index:])
	}

	for index, item := range an_array_or_slice { // 数组或者slice迭代
		block
	}

	for index := range an_array_or_slice { // 数组或者slice迭代
		block // item :=  an_array_or_slice[index]
	}

	for key, value := range aMap { // map迭代
		block
	}

	for key := range aMap { // map迭代
		block // value := aMap[key
	}

	for item := range a_channel { // channel 迭代
		block
	}
```

# 带有标签

```go
func test(table [][]string, x string) {
	found := false
	for row := range table {
		for col := range table[row] {
			if table[row][col] == x {
				found = true
				break
			}
			if found {
				break
			}
		}
	}
}

func test1(table [][]string, x string) {
	var found bool = false
FOUND:
	for row := range table {
		for column := range table[row] {
			if table[row][column] == x {
				found = true
				break FOUND
			}
		}
	}
}

```

> 标签可以作用于for ,switch,select;  break continue goto 都可以使用

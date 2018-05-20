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

package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total: ", f) // Total:  24
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others




挑战# 1:
——使用goroutines和channel来计算阶乘。

挑战# 2:
为什么你想用goroutines和channel来计算阶乘?
——制定你自己的答案，然后把答案发到这个讨论区:https://goo.gl/flGsyX。
——在讨论区读一些其他的答案，看看别人的原因。


*/

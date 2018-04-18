package _2_stringutil

func reverseTwo(s string) string {

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 这说明了一个未导出的函数。
// 可以在同一个包中使用导出的函数。

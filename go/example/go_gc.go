cat main.go 
package main

import (
	"fmt"
	"runtime"
	"time"
)

func print_stats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	print_stats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("操作失败")
		}
		time.Sleep(5 * time.Second)
		print_stats(mem)

	}
	print_stats(mem)

}

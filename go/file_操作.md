
#  读取大文件

```go
package main

import (
    "bufio"
    "log"
    "os"
)

func main() {
    file, err := os.Open("large_file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    const maxScanTokenSize = 64 * 1024 * 1024 // 64MB
    buf := make([]byte, maxScanTokenSize)

    scanner := bufio.NewScanner(file)
    scanner.Buffer(buf, maxScanTokenSize)

    for scanner.Scan() {
        line := scanner.Text()
        // 处理每一行的逻辑
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

```
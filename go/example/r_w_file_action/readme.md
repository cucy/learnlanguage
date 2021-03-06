[toc]

# io pkg

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	// 打开文件
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件错误: %s, %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	buf := make([]byte, 8)
	if _, err := io.ReadFull(f, buf); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}
	io.WriteString(os.Stdout, string(buf))
	fmt.Println()

}

```

# io.Reader io.Writer

```go
package main

import (
	"fmt"
	"io"
	"os"
)

// 实现读方法
func count_chars(r io.Reader) int {
	buf := make([]byte, 16)
	total := 0
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}
		if err == io.EOF {
			break
		}
		total += n
	}
	fmt.Println(total)
	return total
}

func write_number_of_chars(w io.Writer, x int) {
	// 实现了写接口  io.Writer
	fmt.Fprintf(w, "%d\n", x)
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]

	_, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error on file %s: %s\n", filename, err)
		os.Exit(1)
	}
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file:", err)
		os.Exit(-1)
	}
	defer f.Close()

	// 读
	chars := count_chars(f)
	filename = filename + ".count"

	f, err = os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
	}
	defer f.Close()
	// 写
	write_number_of_chars(f, chars)
}

```

# bufio pkg
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening %s: %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	//	 按行读取文件
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Err() != nil {
			fmt.Printf("error reading file %s", err)
			os.Exit(1)
		}
		fmt.Println(line)
	}
}

```



# fmt.Fprintf 

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	destination, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		os.Exit(1)
	}
	defer destination.Close()

	fmt.Fprintf(destination, "[%s]:", filename)
	fmt.Fprintf(destination, "使用 fmt.Fprintf 方法 in %s\n", filename)

}

/*
$ go run fmtF.go test
$ cat test
[test]: Using fmt.Fprintf in test
*/

```

# io.copy

```go 
package main

import (
	"fmt"
	"io"
	"os"
)

func Copy(src, dst string) (int64, error) {
	source_file_stat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !source_file_stat.Mode().IsRegular() {
		// 是否是常规文件
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	// 打开源文件
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	// 创建目标文件
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	n_bytes, err := io.Copy(destination, source)
	return n_bytes, err

}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two command line arguments!")
		os.Exit(1)
	}
	sourceFile := os.Args[1]
	destinationFile := os.Args[2]

	nBytes, err := Copy(sourceFile, destinationFile)

	if err != nil {
		fmt.Printf("The copy operation failed %q\n", err)
	} else {
		fmt.Printf("Copied %d bytes!\n", nBytes)
	}
}

/*
go run main.go src_file des_file
Copied 22 bytes!


*/

```

## 方法2 reading a file all at once

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two command line arguments!")
		os.Exit(1)
	}
	sourceFile := os.Args[1]
	destinationFile := os.Args[2]

	// 读
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 写
	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating the new file", destinationFile)
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
go run   main.go src_file des_file

ioutil 会自动打开关闭文件
*/

```

## 方法3 比较好的方式

```go
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var BUFFERSIZE int64

func Copy(src, dst string, BUFFERSIZE int64) error {
	source_file_stat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !source_file_stat.Mode().IsRegular() {
		//  不是常规文件
		return fmt.Errorf("%s is not a regular file.", src)
	}

	//  打开文件
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// 写文件
	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		// 读取文件
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		// 写入文件
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("usage: %s source destination BUFFERSIZE\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	source := os.Args[1]
	destination := os.Args[2]
	BUFFERSIZE, _ = strconv.ParseInt(os.Args[3], 10, 64)

	fmt.Printf("Copying %s to %s\n", source, destination)
	err := Copy(source, destination, BUFFERSIZE)
	if err != nil {
		fmt.Printf("File copying failed: %q\n", err)
	}
}

/*
go run main.go src_file des_file 1024
Copying src_file to des_file


*/

## io.Open

```


func main() {

	f, err := Readfile(`./textfile.log`)

	if err!=nil{
		fmt.Println(err)
	}

	fmt.Print(string(f))

}

func Readfile(fname string) ([]byte, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 1024) //  一次读取1024字节
	tmp_read := make([]byte, 1024) //  读取结果暂存
FILE_READ:
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break FILE_READ
		}
		tmp_read = append(tmp_read, buf[:n]...)
	}

	return tmp_read, nil

}



```go


```

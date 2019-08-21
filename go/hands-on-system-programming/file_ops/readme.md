# 文件操作


## 获取当前运行目录

```go
// 当前目录
wd, err := os.Getwd()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("starting dir:", wd)

// 切换目录
if err := os.Chdir("/"); err != nil {
    fmt.Println(err)
    return
}

if wd, err = os.Getwd(); err != nil {
    fmt.Println(err)
    return
}
fmt.Println("final dir:", wd)
```

```go
// 返回绝对路径
func Abs(path string) (string, error):  

// 返回basename
func Base(path string) string

// 
func Clean(path string) string: 

// 文件扩展名
func Ext(path string) string:
```

## 读文件

```go
s := strings.NewReader(`Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.`)
    io.Copy(os.Stdout, io.LimitReader(s, 25)) // will print "Lorem Ipsum is simply dum"
```

> 设置颜色

```go

type queryWriter struct {
	Query []byte
	io.Writer
}

func (q queryWriter) Write(b []byte) (n int, err error) {
	lines := bytes.Split(b, []byte{'\n'})
	l := len(q.Query)
	for _, b := range lines {
		i := bytes.Index(b, q.Query)
		if i == -1 {
			continue
		}
		for _, s := range [][]byte{
			b[:i],              // what's before the match
			[]byte("\x1b[31m"), //star red color
			b[i : i+l],         // match
			[]byte("\x1b[39m"), // default color
			b[i+l:],            // whatever is left
		} {
			v, err := q.Writer.Write(s)
			n += v
			if err != nil {
				return 0, err
			}
		}
		fmt.Fprintln(q.Writer)
	}
	return len(b), nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please specify a path and a search string.")
		return
	}
	root, err := filepath.Abs(os.Args[1]) // get absolute path
	if err != nil {
		fmt.Println("Cannot get absolute path:", err)
		return
	}
	query := []byte(strings.Join(os.Args[2:], " "))
	fmt.Printf("Searching for %q in %s...\n", query, root)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = ioutil.ReadAll(io.TeeReader(f, queryWriter{Query: query, Writer: os.Stdout}))
		return err
	})
	if err != nil {
		fmt.Println(err)
	}

}
```



```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please specify a path.")
        return
    }
    b, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(string(b))
}
```

> 自定义缓冲大小读取文件

```go
func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please specify a file")
        return
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close() // we ensure close to avoid leaks

    var (
        b = make([]byte, 16)
    )
    for n := 0; err == nil; {
        n, err = f.Read(b)
        if err == nil {
            fmt.Print(string(b[:n])) // only print what's been read
        }
    }
    if err != nil && err != io.EOF { // we expect an EOF
        fmt.Println("\n\nError:", err)
    }
}
```


> 缓冲

```go
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var b = bytes.NewBuffer(make([]byte, 26))
    var texts = []string{
        `As he came into the window`,
        `It was the sound of a crescendo
He came into her apartment`,
        `He left the bloodstains on the carpet`,
        `She ran underneath the table
He could see she was unable
So she ran into the bedroom
She was struck down, it was her doom`,
    }
    for i := range texts {
        b.Reset()
        b.WriteString(texts[i])
        fmt.Println("Length:", b.Len(), "\tCapacity:", b.Cap())
    }
}
```

> 统计文件行数

```go
func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please specify a path.")
        return
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close()
    r := bufio.NewReader(f) // wrapping the reader with a buffered one
    var rowCount int
    for err == nil {
        var b []byte
        for moar := true; err == nil && moar; {
            b, moar, err = r.ReadLine()
            if err == nil {
                fmt.Print(string(b))
            }
        }
        // each time moar is false, a line is completely read
        if err == nil {
            fmt.Println()
            rowCount++

        }
    }
    if err != nil && err != io.EOF {
        fmt.Println("\nError:", err)
        return
    }
    fmt.Println("\nRow count:", rowCount)
}
```

## 写文件

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Please specify a path and some content")
        return
    }
    // the second argument, the content, needs to be casted to a byte slice
    if err := ioutil.WriteFile(os.Args[1], []byte(os.Args[2]), 0644); err != nil {
        fmt.Println("Error:", err)
    }
}
```


> 反转文件内容


```go
// Let's omit argument check and file opening, we obtain src and dst
cur, err := src.Seek(0, os.SEEK_END) // Let's go to the end of the file
if err != nil {
    fmt.Println("Error:", err)
    return
}
b := make([]byte, 16)

// 在移动到文件末尾并定义字节缓冲区之后，我们在文件中输入一个稍微向后移动的循环，然后读取其中的一部分，如下面的代码所示：
for step, r, w := int64(16), 0, 0; cur != 0; {
    if cur < step { // ensure cursor is 0 at max
        b, step = b[:cur], cur
    }
    cur = cur - step
    _, err = src.Seek(cur, os.SEEK_SET) // go backwards
    if err != nil {
        break
    }
    if r, err = src.Read(b); err != nil || r != len(b) {
        if err == nil { // all buffer should be read
            err = fmt.Errorf("read: expected %d bytes, got %d", len(b), r)
        }
        break
    }
}

// 然后，我们反转内容并将其写入目的地，如下代码所示：

for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        switch { // Swap (\r\n) so they get back in place
        case b[i] == '\r' && b[i+1] == '\n':
            b[i], b[i+1] = b[i+1], b[i]
        case j != len(b)-1 && b[j-1] == '\r' && b[j] == '\n':
            b[j], b[j-1] = b[j-1], b[j]
            }
            b[i], b[j] = b[j], b[i] // swap bytes
        }
        if w, err = dst.Write(b); err != nil || w != len(b) {
            if err != nil {
                err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
            }
        }
}
    if err != nil && err != io.EOF { // we expect an EOF
        fmt.Println("\n\nError:", err)
    }
}
```

> 缓冲写


```go
const grr = "G.R.R. Martin"

type book struct {
    Author, Title string
    Year int
}

func main() {
    dst, err := os.OpenFile("book_list.txt", os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer dst.Close()
    bookList := []book{
        {Author: grr, Title: "A Game of Thrones", Year: 1996},
        {Author: grr, Title: "A Clash of Kings", Year: 1998},
        {Author: grr, Title: "A Storm of Swords", Year: 2000},
        {Author: grr, Title: "A Feast for Crows", Year: 2005},
        {Author: grr, Title: "A Dance with Dragons", Year: 2011},
        // if year is omitted it defaulting to zero value
        {Author: grr, Title: "The Winds of Winter"},
        {Author: grr, Title: "A Dream of Spring"},
    }
    b := bytes.NewBuffer(make([]byte, 0, 16))
    for _, v := range bookList {
        // prints a msg formatted with arguments to writer
        fmt.Fprintf(b, "%s - %s", v.Title, v.Author)
        if v.Year > 0 { 
            // we do not print the year if it's not there
            fmt.Fprintf(b, " (%d)", v.Year)
        }
        b.WriteRune('\n')
        if _, err := b.WriteTo(dst); true { // copies bytes, drains buffer
            fmt.Println("Error:", err)
            return
        }
    }
}
```



```go
	r := strings.NewReader("let's read this message\n")
	b := bytes.NewBuffer(nil)
	w := io.MultiWriter(b, os.Stdout)
	io.Copy(w, r)
	fmt.Println(b.String())
```

## 操作文件第三方包

https://github.com/blang/vfs

https://github.com/spf13/afero


# 文件流操作

```go
func CopyNOffset(dst io.Writer, src io.ReadSeeker, offset, length int64) (int64, error) {
  if _, err := src.Seek(offset, io.SeekStart); err != nil {
    return 0, err
  }
  return io.CopyN(dst, src, length)
}
```
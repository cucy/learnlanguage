## 标准输入
程序可能需要知道的前两件事是其标识符和父标识符，即PID和PPID。这实际上非常简单–os.getpid（）和os.getppid（）函数都返回带有两个标识符的整数值，如下代码所示：

```go
  fmt.Println("Current PID:", os.Getpid())
  fmt.Println("Current Parent PID:", os.Getppid())
```


# 子进程

```go
Go应用程序可以与操作系统交互以创建一些其他进程。操作系统的另一个子包提供了创建和运行新进程的功能。在os/exec包中，有一个cmd类型，它表示命令执行：
```

```go
 

type Cmd struct {
    Path string // command to run.
    Args []string // command line arguments (including command)
    Env []string // environment of the process
    Dir string // working directory
    Stdin io.Reader // standard input`
    Stdout io.Writer // standard output
    Stderr io.Writer // standard error
    ExtraFiles []*os.File // additional open files
    SysProcAttr *syscall.SysProcAttr // os specific attributes
    Process *os.Process // underlying process
    ProcessState *os.ProcessState // information on exited processte
}
```

创建新命令的最简单方法是使用exec.command函数，该函数采用可执行路径和一系列参数。让我们来看一个带有echo命令和一些参数的简单示例：


```go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("echo", "A", "sample", "command")
    fmt.Println(cmd.Path, cmd.Args[1:]) // echo [A sample command]
}
```


一个非常重要的细节是标准输入、输出和错误的性质——它们都是我们已经熟悉的接口：




```go
The input is an io.Reader, which could be bytes.Reader, bytes.Buffer, strings.Reader, os.File, or any other implementation.
The output and the error are io.Writer, can also be os.File or bytes.Buffer, and can also be strings.Builder or any another writer implementation.
There are different ways to launch the process, depending on what the parent application needs:

Cmd.Run: Executes the command, and returns an error that is nil if the child process is executed correctly.
Cmd.Start : Executes the command asynchronously and lets the parent continue its flow. In order to wait for the child process to finish its execution, there is another method, Cmd.Wait.
Cmd.Output: Executes the command and returns its standard output, and returns an error if Stderr isn't defined but the standard error produced the output.
Cmd.CombinedOutput: Executes the command and returns both a standard error and output combined, which is very useful when the entire output of the child process-standard output plus standard error needs to be checked or saved.
```

```go
输入是一个IO.reader，可以是bytes.reader、bytes.buffer、strings.reader、os.file或任何其他实现。
输出和错误是io.writer，也可以是os.file或bytes.buffer，还可以是strings.builder或任何其他编写器实现。
根据父应用程序的需要，有不同的方法启动流程：
run：执行命令，如果子进程执行正确，则返回一个为零的错误。
start：异步执行命令，并让父级继续其流。为了等待子进程完成其执行，还有另一个方法cmd.wait。
output：执行命令并返回其标准输出，如果未定义stderr，但标准错误生成了输出，则返回错误。
cmd.combined output：执行命令并返回标准错误和输出组合，这在需要检查或保存子进程标准输出加上标准错误的整个输出时非常有用。
```


```go
    cmd := exec.Command("ls", "-l")
    if err := cmd.Start(); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Cmd: ", cmd.Args[0])
    fmt.Println("Args:", cmd.Args[1:])
    fmt.Println("PID: ", cmd.Process.Pid)
    cmd.Wait()
```


> 标准输入

```go
    b := bytes.NewBuffer(nil)
    cmd := exec.Command("cat")
    cmd.Stdin = b
    cmd.Stdout = os.Stdout
    fmt.Fprintf(b, "Hello World! I'm using this memory address: %p", b)
    if err := cmd.Start(); err != nil {
        fmt.Println(err)
        return
    }
    cmd.Wait()
```


https://github.com/kardianos/service


https://github.com/takama/daemon
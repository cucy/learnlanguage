
标准输入可用于接收用户命令。我们可以从使用缓冲输入来读取行并打印它们开始。为了读取一行，有一个有用的命令`bufio.scanner`，它已经提供了一个行阅读器。代码将类似于以下代码段：


由于此代码没有退出点，我们可以从创建第一个命令exit开始，该命令将终止shell的执行。我们可以在代码中做一个小的更改来实现这一点，如下所示：

```go
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Fprint(w, "欢迎使用Sehll\n")
	for {
		s.Scan() // get   next the token
		msg := string(s.Bytes())
		if msg == "exit" {
			return
		}
		fmt.Fprintf(w, "Your wrote %q\n", msg)
	}
```

为了能够正确地解释命令，需要将消息拆分为参数。这与操作系统应用于传递给进程的参数的逻辑相同。strings.split函数通过将空格指定为第二个参数并将字符串拆分为单词来实现此技巧，如下代码所示：


```go
args := strings.Split(string(s.Bytes()), " ")
cmd := args[0]
args = args[1:]
```

```go
switch cmd {
case "exit":
    return
case "someCommand":
    someCommand(w, args)
case "anotherCommand":
    anotherCommand(w, args)
}
```

# Command execution

既然一切都设置好了，那么只剩下定义各种命令的实际作用了。我们可以定义执行命令的函数类型以及开关的行为：

```go
var cmdFunc func(w io.Writer, args []string) (exit bool)
switch cmd {
case "exit":
    cmdFunc = exitCmd
}
if cmdFunc == nil {
    fmt.Fprintf(w, "%q not found\n", cmd)
    continue
}
if cmdFunc(w, args) { // execute and exit if true
    return
}
```

返回值告诉应用程序是否需要终止，并允许我们轻松定义出口函数，而不是特殊情况：

```go
func exitCmd(w io.Writer, args []string) bool {
    fmt.Fprintf(w, "Goodbye! :)")
    return true
}
```


根据应用程序的范围，我们现在可以实现任何类型的命令。让我们创建一个shuffle命令，它将在math/rand包的帮助下按无序顺序打印参数：

```go
func shuffle(w io.Writer, args ...string) bool {
    rand.Shuffle(len(args), func(i, j int) {
        args[i], args[j] = args[j], args[i]
    })
    for i := range args {
        if i > 0 {
            fmt.Fprint(w, " ")
        }
        fmt.Fprintf(w, "%s", args[i])
    }
    fmt.Fprintln(w)
    return false
}
```


我们可以通过创建一个print命令与文件系统和文件进行交互，该命令将显示在文件内容的输出中：

```go
func print(w io.Writer, args ...string) bool {
    if len(args) != 1 {
        fmt.Fprintln(w, "Please specify one file!")
        return false
    }
    f, err := os.Open(args[0])
    if err != nil {
        fmt.Fprintf(w, "Cannot open %s: %s\n", args[0], err)
    }
    defer f.Close()
    if _, err := io.Copy(w, f); err != nil {
        fmt.Fprintf(w, "Cannot print %s: %s\n", args[0], err)
    }
    fmt.Fprintln(w)
    return false
}
```




## Some refactor 重构

只要稍加重构，就可以改进当前版本的伪终端应用程序。我们可以从将命令定义为自定义类型开始，并使用一些描述其行为的方法：

```go
type cmd struct {
	Name   string // the command name
	Help   string // a description string
	Action func(w io.Writer, args ...string) bool
}

func (c cmd) Match(s string) bool {
	return c.Name == s
}
func (c cmd) Run(w io.Writer, args ...string) bool {
	return c.Action(w, args...)
}
```

每个命令的所有信息都可以独立于一个结构中。我们还可以开始定义依赖于其他命令的命令，例如帮助命令。如果在var cmds[]cmd包中的某个地方定义了命令切片或映射，那么help命令将如下所示：

```go
help = cmd{
    Name: "help", Help: "Shows available commands",
	Action: func(w io.Writer, args ...string) bool {
		fmt.Fprintln(w, "Available commands:")
		for _, c := range cmds {
			fmt.Fprintf(w, " - %-15s %s\n", c.Name, c.Help)
		}
		return false
	},}
```


主循环中选择正确命令的部分将略有不同；它需要在切片中找到匹配项并执行它：


```go
for i := range cmds {
    if !cmds[i].Match(args[0]) {
        continue
    }
    idx = i
    break
}

if idx == -1 {
    fmt.Fprintf(w, "%q not found. Use `help` for available commands\n", args[0])
    continue
}

if cmds[idx].Run(w, args[1:]...) {
    fmt.Fprintln(w)
    return
}
```
现在有一个显示可用命令列表的帮助命令，我们可以主张每次用户指定一个不存在的命令时使用它，因为我们当前正在检查索引是否已从其默认值-1更改。


## Multiline input

可以改进的第一件事是通过添加对带引号字符串的支持来改善参数和间距之间的关系。这可以通过bufio.scanner实现，它有一个自定义的split函数，其行为类似于bufio.scanwords，除了它知道引号之外。下面的代码演示了这一点：

```go
func ScanArgs(data []byte, atEOF bool) (advance int, token []byte, err error) {
    // first space
    start, first := 0, rune(0)
    for width := 0; start < len(data); start += width {
        first, width = utf8.DecodeRune(data[start:])
        if !unicode.IsSpace(first) {
            break
        }
    }
    // skip quote
    if isQuote(first) {
        start++
    }
```



函数有一个第一个块，该块跳过空格并查找第一个非空格字符；如果该字符是引号，则跳过它。然后，它查找终止参数的第一个字符，该字符是正常参数的空格，以及其他参数的相应引号：

```go
    // loop until arg end character
    for width, i := 0, start; i < len(data); i += width {
        var r rune
        r, width = utf8.DecodeRune(data[i:])
        if ok := isQuote(first); !ok && unicode.IsSpace(r) || ok  
            && r == first {
                return i + width, data[start:i], nil
        }
    }
```


如果在引用上下文中到达文件结尾，则返回部分字符串；否则，不会跳过引用并请求更多数据：

```go
    // token from EOF
    if atEOF && len(data) > start {
        return len(data), data[start:], nil
    }
    if isQuote(first) {
        start--
    }
    return start, nil, nil
}
```

> 完全例子

```go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := bufio.NewScanner(strings.NewReader(`And the 'cats in 
	the craddle' and "the silver's spoon" nasa put a men on "the moon`))
	s.Split(ScanArgs)
	for s.Scan() {
		fmt.Printf("%q\n", s.Text())
	}
	fmt.Println(s.Err())
}

var ErrClosingQuote = errors.New("Missing closing quote")

func isQuote(r rune) bool {
	return r == '"' || r == '\''
}

func ScanArgs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start, first := 0, rune(0)
	for width := 0; start < len(data); start += width {
		first, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(first) {
			break
		}
	}
	if isQuote(first) {
		start++
	}
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if ok := isQuote(first); !ok && unicode.IsSpace(r) || ok && r == first {
			return i + width, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		if isQuote(first) {
			err = ErrClosingQuote
		}
		return len(data), data[start:], err
	}
	if isQuote(first) {
		start--
	}
	return start, nil, nil
}

```



现在我们可以使用这一行来解析参数，同时使用helper结构argscanner，定义如下：

```go
type argsScanner []string

func (a *argsScanner) Reset() { *a = (*a)[0:0] }

func (a *argsScanner) Parse(r io.Reader) (extra string) {
    s := bufio.NewScanner(r)
    s.Split(ScanArgs)
    for s.Scan() {
        *a = append(*a, s.Text())
    }
    if len(*a) == 0 {
        return ""
    }
    lastArg := (*a)[len(*a)-1]
    if !isQuote(rune(lastArg[0])) {
        return ""
    }
    *a = (*a)[:len(*a)-1]
    return lastArg + "\n"
}
```
此自定义切片将允许我们通过更改循环的工作方式来接收带引号的行和引号之间的新行：

```go
func main() {
 s := bufio.NewScanner(os.Stdin)
 w := os.Stdout
 a := argsScanner{}
 b := bytes.Buffer{}
 for {
        // prompt message  提示
        a.Reset()
        b.Reset()
        for {
            s.Scan()
            b.Write(s.Bytes())
            extra := a.Parse(&b)
            if extra == "" {
                break
            }
            b.WriteString(extra)
        }
        // a contains the split arguments
    }
}
```

## Providing color support to the pseudo-terminal 提供颜色

伪终端可以通过提供彩色输出来改进。我们已经看到，在Unix中，有转义序列可以更改背景和前景颜色。我们先定义一个自定义类型：

```go
type color int

func (c color) Start(w io.Writer) {
    fmt.Fprintf(w, "\x1b[%dm", c)
}

func (c color) End(w io.Writer) {
    fmt.Fprintf(w, "\x1b[%dm", Reset)
}

func (c color) Sprintf(w io.Writer, format string, args ...interface{}) {
    c.Start(w)
    fmt.Fprintf(w, format, args...)
    c.End(w)
}

// List of colors
const (
    Reset color = 0
    Red color = 31
    Green color = 32
    Yellow color = 33
    Blue color = 34
    Magenta color = 35
    Cyan color = 36
    White color = 37
)
```

这种新类型可以用来增强彩色输出的命令。例如，既然我们支持带空格的参数，那么让我们使用shuffle命令并使用替换颜色来区分字符串：


```go
func shuffle(w io.Writer, args ...string) bool {
    rand.Shuffle(len(args), func(i, j int) {
        args[i], args[j] = args[j], args[i]
    })
    for i := range args {
        if i > 0 {
            fmt.Fprint(w, " ")
        }
        var f func(w io.Writer, format string, args ...interface{})
        if i%2 == 0 {
            f = Red.Fprintf
        } else {
            f = Green.Fprintf
        }
        f(w, "%s", args[i])
    }
    fmt.Fprintln(w)
    return false
}
```

## Suggesting commands  建议命令

当指定的命令不存在时，我们可以建议使用一些类似的命令。为此，我们可以使用Levenshtein距离公式，该公式通过计算从一个字符串到另一个字符串所需的删除、插入和替换来测量字符串之间的相似性。
在下面的代码中，我们将使用agnivade/levenshtein包，该包将通过go get命令获得：

```go
go get github.com/agnivade/levenshtein/...
```


Then, we define a new function to call when there is no match with existing commands:

```go
func commandNotFound(w io.Writer, cmd string) {
    var list []string
    for _, c := range cmds {
        d := levenshtein.ComputeDistance(c.Name, cmd)
        if d < 3 {
            list = append(list, c.Name)
        }
    }
    fmt.Fprintf(w, "Command %q not found.", cmd)
    if len(list) == 0 {
        return
    }
    fmt.Fprint(w, " Maybe you meant: ")
    for i := range list {
        if i > 0 {
            fmt.Fprint(w, ", ")
        }
        fmt.Fprintf(w, "%s", list[i])
    }
}
```

## Extensible commands 可扩展命令


目前，伪终端的局限性在于它的可扩展性。如果需要添加新命令，则需要直接将其添加到主包中。我们可以考虑将命令与主包分离，并允许其他用户使用其命令扩展功能：

1 第一步是创建导出的命令。让我们使用一个接口来定义一个命令，这样用户就可以实现自己的命令：

```go
// Command represents a terminal command
type Command interface {
    GetName() string
    GetHelp() string
    Run(input io.Reader, output io.Writer, args ...string) (exit bool)
}
```

2 现在，我们可以为其他包指定命令列表和函数来添加其他命令：

```go
// ErrDuplicateCommand is returned when two commands have the same name
var ErrDuplicateCommand = errors.New("Duplicate command")

var commands []Command

// Register adds the Command to the command list
func Register(command Command) error {
    name := command.GetName()
    for i, c := range commands {
        // unique commands in alphabetical order
        switch strings.Compare(c.GetName(), name) {
        case 0:
            return ErrDuplicateCommand
        case 1:
            commands = append(commands, nil)
            copy(commands[i+1:], commands[i:])
            commands[i] = command
            return nil
        case -1:
            continue
        }
    }
    commands = append(commands, command)
    return nil
}
```

3.我们可以提供命令的基本实现，以执行简单的函数：

```go
// Base is a basic Command that runs a closure
type Base struct {
    Name, Help string
    Action func(input io.Reader, output io.Writer, args ...string) bool
}

func (b Base) String() string { return b.Name }

// GetName returns the Name
func (b Base) GetName() string { return b.Name }

// GetHelp returns the Help
func (b Base) GetHelp() string { return b.Help }

// Run calls the closure
func (b Base) Run(input io.Reader, output io.Writer, args ...string) bool {
    return b.Action(input, output, args...)
}
```

4 我们可以提供一个与命令名匹配的函数：

```go
// GetCommand returns the command with the given name
func GetCommand(name string) Command {
    for _, c := range commands {
        if c.GetName() == name {
            return c
        }
    }
    return suggest
}
```



5 我们可以使用前一个示例中的逻辑使该函数返回suggestion命令，其定义如下：

```go
var suggest = Base{
    Action: func(in io.Reader, w io.Writer, args ...string) bool {
        var list []string
        for _, c := range commands {
            name := c.GetName()
            d := levenshtein.ComputeDistance(name, args[0])
            if d < 3 {
                list = append(list, name)
            }
        }
        fmt.Fprintf(w, "Command %q not found.", args[0])
        if len(list) == 0 {
            return false
        }
        fmt.Fprint(w, " Maybe you meant: ")
        for i := range list {
            if i > 0 {
                fmt.Fprint(w, ", ")
            }
            fmt.Fprintf(w, "%s", list[i])
        }
        return false
    },
}
```

6 我们已经可以在出口和帮助包中注册几个命令。此处只能定义帮助，因为命令列表是私有的：

```go
func init() {
    Register(Base{Name: "help", Help: "...", Action: helpAction})
    Register(Base{Name: "exit", Help: "...", Action: exitAction})
}

func helpAction(in io.Reader, w io.Writer, args ...string) bool {
    fmt.Fprintln(w, "Available commands:")
    for _, c := range commands {
        n := c.GetName()
        fmt.Fprintf(w, " - %-15s %s\n", n, c.GetHelp())
    }
    return false
}

func exitAction(in io.Reader, w io.Writer, args ...string) bool {
    fmt.Fprintf(w, "Goodbye! :)\n")
    return true
}
```

此方法将允许用户使用commandbase结构创建简单命令，或者在其命令需要时嵌入该命令或使用自定义结构（就像具有状态的命令）：
This approach will allow a user to use the commandBase struct to create a simple command, or to embed it or use a custom struct if their command requires it (like a command with a state):



```go
// Embedded unnamed field (inherits method)
type MyCmd struct {
    Base
    MyField string
}

// custom implementation
type MyImpl struct{}

func (MyImpl) GetName() string { return "myimpl" }
func (MyImpl) GetHelp() string { return "help string"}
func (MyImpl) Run(input io.Reader, output io.Writer, args ...string) bool {
    // do something
    return true
}
```
mycmd结构和myimpl结构的区别在于，一个可以用作另一个命令的修饰器，而另一个则是不同的实现，因此它不能与另一个命令交互。





## Commands with status 具有状态的命令

到目前为止，我们已经创建了没有内部状态的命令。但有些命令可以保持内部状态并相应地更改其行为。状态可以限制在会话本身，也可以在多个会话之间共享。更明显的例子是终端中的命令历史记录，在那里执行的所有命令都在会话之间存储和保留。


## Volatile status 挥发性状态

最容易实现的是一种状态，它不是持久的，当应用程序退出时会丢失。我们所需要做的就是创建一个定制的数据结构，它承载状态并满足命令接口。这些方法将属于指向该类型的指针，否则它们将无法修改数据。
在下面的示例中，我们将创建一个非常基本的内存存储，它作为带有参数的堆栈（先进、后出）工作。让我们从推送和弹出功能开始：

```go
type Stack struct {
    data []string
}

func (s *Stack) push(values ...string) {
    s.data = append(s.data, values...)
}

func (s *Stack) pop() (string, bool) {
    if len(s.data) == 0 {
        return "", false
    }
    v := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return v, true
}
```

存储在堆栈中的字符串表示命令的状态。现在，我们需要实现命令接口的方法，我们可以从最简单的方法开始：

```go
func (s *Stack) GetName() string {
    return "stack"
}

func (s *Stack) GetHelp() string {
    return "a stack-like memory storage"
}
```


现在我们需要决定它如何在内部工作。将有两个子命令：
push后面跟着一个或多个参数，将被推送到堆栈。
pop将采用堆栈的最顶层元素，它不需要任何参数。
让我们定义一个helper方法isvalid，它检查参数是否有效：

```go
func (s *Stack) isValid(cmd string, args []string) bool {
    switch cmd {
    case "pop":
        return len(args) == 0
    case "push":
        return len(args) > 0
    default:
        return false
    }
}
```

现在，我们可以实现将使用有效性检查的命令执行方法。如果通过，它将执行所选命令或显示帮助消息：

```go
func (s *Stack) Run(r io.Reader, w io.Writer, args ...string) (exit bool) {
    if l := len(args); l < 2 || !s.isValid(args[1], args[2:]) {
        fmt.Fprintf(w, "Use `stack push <something>` or `stack pop`\n")
        return false
    }
    if args[1] == "push" {
        s.push(args[2:]...)
        return false
    }
    if v, ok := s.pop(); !ok {
        fmt.Fprintf(w, "Empty!\n")
    } else {
        fmt.Fprintf(w, "Got: `%s`\n", v)
    }
    return false
}
```


## Persistent status  持续状态          
                                             
下一步是在会话之间保持状态，这要求在应用程序启动时执行一些操作，在应用程序结束时执行另一个操作。这些新行为可以与命令界面上的一些更改集成在一起：

```go
type Command interface {
    Startup() error
    Shutdown() error
    GetName() string
    GetHelp() string
    Run(r io.Reader, w io.Writer, args ...string) (exit bool)
}
```

startup（）方法负责应用程序启动时的状态和加载，shutdown（）方法需要在退出前将当前状态保存到磁盘。我们可以用这些方法更新基础结构；但是，这不会起任何作用，因为没有状态：

```go
// Startup does nothing
func (b Base) Startup() error { return nil }

// Shutdown does nothing
func (b Base) Shutdown() error { return nil }
```

不导出命令列表；它是未导出的变量commands。我们可以添加两个与此类列表交互的函数，并确保对所有可用命令执行以下方法：启动和关闭：

```go
// Shutdown executes shutdown for all commands
func Shutdown(w io.Writer) {
    for _, c := range commands {
        if err := c.Shutdown(); err != nil {
            fmt.Fprintf(w, "%s: shutdown error: %s", c.GetName(), err)
        }
    }
}

// Startup executes Startup for all commands
func Startup(w io.Writer) {
    for _, c := range commands {
        if err := c.Startup(); err != nil {
            fmt.Fprintf(w, "%s: startup error: %s", c.GetName(), err)
        }
    }
}
```


最后一步是在启动主循环之前在主应用程序中使用这些函数：

```go
func main() {
    s, w, a, b := bufio.NewScanner(os.Stdin), os.Stdout, args{}, bytes.Buffer{}
    command.Startup(w)
    defer command.Shutdown(w) // this is executed before returning
    fmt.Fprint(w, "** Welcome to PseudoTerm! **\nPlease enter a command.\n")
    for {
        // main loop
    }
}
```

## Upgrading the Stack command 升级stack命令

我们希望之前定义的命令stack能够在会话之间保存其状态。最简单的解决方案是将堆栈的内容保存为文本文件，每行一个元素。我们可以使用操作系统/用户包将此文件放置在用户主目录中，从而使每个用户的文件都是唯一的：

```go
func (s *Stack) getPath() (string, error) {
    u, err := user.Current()
    if err != nil {
        return "", err
    }
    return filepath.Join(u.HomeDir, ".stack"), nil
}
```


让我们开始编写；我们将创建并截断文件（使用trunc标志将其大小设置为0），并编写以下行：

```go
func (s *Stack) Shutdown(w io.Writer) error {
    path, err := s.getPath()
    if err != nil {
        return err
    }
    f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
    if err != nil {
        return err
    }
    defer f.Close()
    for _, v := range s.data {
        if _, err := fmt.Fprintln(f, v); err != nil {
            return err
        }
    }
    return nil
}
```

关机期间使用的方法将逐行读取文件，并将元素添加到堆栈中。我们可以使用bufio.scanner，正如我们在前一章中看到的那样，很容易做到这一点：

```go
func (s *Stack) Startup(w io.Writer) error {
    path, err := s.getPath()
    if err != nil {
        return err
    }
    f, err := os.Open(path)
    if err != nil {
        if os.IsNotExist(err) {
            return nil
        }
        return err
    }
    defer f.Close()
    s.data = s.data[:0]
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        s.push(string(scanner.Bytes()))
    }
    return nil
}
```

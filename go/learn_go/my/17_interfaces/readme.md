!(interface)[https://raw.githubusercontent.com/cucy/learnlanguage/master/go/learn_go/my/17_interfaces/img/interfacedfds.jpg]


"Polymorphism is the ability to write code that can take on different behavior through the
 implementation of types. Once a type implements an interface, an entire world of
 functionality can be opened up to values of that type."
 - Bill Kennedy

"Interfaces are types that just declare behavior. This behavior is never implemented by the
 interface type directly, but instead by user-defined types via methods. When a
 user-defined type implements the set of methods declared by an interface type, values of
 the user-defined type can be assigned to values of the interface type. This assignment
 stores the value of the user-defined type into the interface value.

 If a method call is made against an interface value, the equivalent method for the
 stored user-defined value is executed. Since any user-defined type can implement any
 interface, method calls against an interface value are polymorphic in nature. The
 user-defined type in this relationship is often called a concrete type, since interface values
 have no concrete behavior without the implementation of the stored user-defined value."
  - Bill Kennedy

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

Values          Receivers
-----------------------------------------------
T               (t T)
*T              (t T) and (t *T)


SOURCE:
Go In Action
William Kennedy
/////////////////////////////////////////////////////////////////////////

Interface types express generalizations or abstractions about the behaviors of other types.
By generalizing, interfaces let us write functions that are more flexible and adaptable
because they are not tied to the details of one particular implementation.

Many object-oriented lagnuages have some notion of interfaces, but what makes Go's interfaces
so distinctive is that they are SATISIFIED IMPLICITLY. In other words, there's no need to declare
all the interfaces that a given CONCRETE TYPE satisifies; simply possessing the necessary methods
is enough. This design lets you create new interfaces that are satisifed by existing CONCRETE TYPES
without changing the existing types, which is particularly useful for types defined in packages that
you don't control.

All the types we've looked at so far have been CONCRETE TYPES. A CONCRETE TYPE specifies the exact
representation of its values and exposes the intrinsic operations of that representation, such as
arithmetic for numbers, or indexing, append, and range for slices. A CONCRETE TYPE may also provide
additional behaviors through its methods. When you have a value of a CONCRETE TYPE, you know exactly
what is IS and what you can DO with it.

There is another kind of type in Go called an INTERFACE TYPE. An interface is an ABSTRACT TYPE. It doesn't
expose the representation or internal structure of its values, or the set of basic operations they support;
it reveals only some of their methods. When you have a value of an interface type, you know nothing about
what it IS; you know only what it can DO, or more precisely, what BEHAVIORS ARE PROVIDED BY ITS METHODS.

-------------------

type ReadWriter interface {
    Reader
    Writer
}

This is called EMBEDDING an interface.


-------------------

A type SATISFIES an interface if it possesses all the methods the interface requires.

-------------------

Conceptually, a value of an interface type, or INTERFACE VALUE, has two components,
    a CONCRETE TYPE and a
    VALUE OF THAT TYPE.
These are called the interface's
    DYNAMIC TYPE and
    DYNAMIC VALUE.

For a statically typed language like Go, types are a compile-time concept, so a type is not a value.
In our conceptual model, a set of values called TYPE DESCRIPTORS provide information about each type,
such as its name and methods. In an interface value, the type component is represented by the appropriate
type descriptor.


var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil


var w io.Writer
w
type: nil
value: nil

w = os.Stdout
w
type: *os.File
value: the address where a value of type os.File is stored

w = new(bytes.Buffer)
w
type: *bytes.Buffer
value: the address where a value of type bytes.Buffer is stored

w = nil
w
type: nil
value: nil

-------------------
The Go Programming Language
Donovan and Kernighan

Caplitalization and identation mine.


------

“多态性是一种编写代码的能力，可以在不同的行为中采取不同的行为。”
类型的实现。一旦一个类型实现了一个接口，一个完整的世界。
功能可以向该类型的值打开。
——比尔•肯尼迪

接口是只声明行为的类型。这种行为从来没有被执行过。
接口类型是直接的，而不是由用户定义的类型通过方法。当一个
用户定义的类型实现了由接口类型、值所声明的方法集。
用户定义的类型可以分配给接口类型的值。这个任务
将用户定义类型的值存储到接口值中。

如果方法调用是针对接口值进行的，则对应的方法是。
将执行存储用户定义的值。因为任何用户定义的类型都可以实现任何类型。
接口、方法对接口值的调用在本质上是多态的。的
这种关系中的用户定义类型通常称为具体类型，因为接口值。
没有实现存储用户定义值的具体行为。
——比尔•肯尼迪

接收器的价值观
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
(t) t和* t。
t(t * t)*

值接收器
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
T(T)
*T (T)和(T *T)


来源:
去行动
威廉•肯尼迪
/ / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / / /

接口类型表示对其他类型的行为的概括或抽象。
通过一般化，接口让我们编写更灵活和适应性更强的函数。
因为它们与一个特定实现的细节无关。

许多面向对象的lagnuages都有一些接口的概念，但是什么使Go的接口成为可能。
它们的独特之处在于它们被含蓄地讽刺了。换句话说，没有必要声明。
一个给定的具体类型满足的所有接口;只是拥有必要的方法。
就足够了。该设计允许您创建满足现有具体类型的新接口。
不改变现有类型，这对于在包中定义的类型特别有用。
你不控制。

到目前为止，我们看到的所有类型都是具体类型。具体类型指定了具体的类型。
表示其值并公开该表示的内部操作，例如。
数字的算术，或索引，附加，和范围的切片。具体类型也可以提供。
通过其方法的附加行为。当你有一个具体类型的值时，你就知道了。
这是什么，你能用它做什么。

还有一种类型叫做接口类型。接口是一种抽象类型。它不
公开其值的表示或内部结构，或它们支持的基本操作集;
它只揭示了他们的一些方法。当您有一个接口类型的值时，您不知道什么。
它是什么;你只知道它能做什么，或者更确切地说，它的方法提供了什么行为。

- - - - - - - - - - - - - - - - - - -

{读写接口类型
读者
作家
}

这称为嵌入接口。


- - - - - - - - - - - - - - - - - - -

一个类型满足接口，如果它拥有接口所要求的所有方法。

- - - - - - - - - - - - - - - - - - -

从概念上讲，接口类型或接口值的值有两个组件，
一个具体的类型和一个。
值的类型。
这些被称为接口。
动态类型和
动态值。

对于像Go这样的静态类型语言，类型是编译时概念，所以类型不是值。
在我们的概念模型中，一组称为类型描述符的值提供关于每种类型的信息，
比如它的名字和方法。在接口值中，类型组件由适当的表示。
类型描述符。


var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil


var w io.Writer
w
类型:零
价值:零

w = os.Stdout
w
类型:* os.File
值:类型操作系统的值的地址。文件存储

w = new(bytes.Buffer)
w
类型:* bytes.Buffer
值:类型字节的值的地址。缓冲存储

w = nil
w
类型:零
价值:零

- - - - - - - - - - - - - - - - - - -
的编程语言
多诺万,克尼汉

我Caplitalization和凹痕。



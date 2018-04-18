 We used shorthand notation:
 to create a variable named p1 of type person
 to create a variable named p2 of type person
 We initialized those variables with specific values
 We used the short variable declaration operator with a struct literal to initialize
 ----------------------------------------
 here is how we talk about structs:
 -- user defined type
 -- we declare the type
 -- the type has fields
 -- the type can also have "tags"
 ----  we haven't seen this yet
 -- the type has an underlying type
 ---- in this case, the underlying type is struct
 -- we declare variables of the type
 -- we initialize those variables
 ---- initialize with a specific value, or
 ---- or, initiliaze to the zero value
 -- a struct is a composite type
 ----------------------------------------
 Bill Kennedy:
 Go allows us the ability to declare our own types.
 Struct types are declared by composing a fixed set of unique fields together.
 Each field in a struct is declared with a known type.
 This could be a built-in type or another user defined type.
 Once we have a type declared, we can create values from the type
 When we declare variables, the value that the variable represents is always initialized.
 The value can be initialized with a specific value or it can be initialized to its zero value
 For numeric types, the zero value would be 0; for strings it would be empty;
 and for booleans it would be false.
 In the case of a struct, the zero value would apply to all the different fields in the struct.
 Anytime a variable is created and initialized to its zero value, it is idiomatic to use the keyword var.
 Reserve the use of the keyword var as a way to indicate that a variable is being set to its zero value.
 If the variable will be initialized to something other than its zero value,
 then use the short variable declaration operator with a struct literal





我们使用速记符号:
创建一个名为p1的变量。
要创建一个名为p2类型的变量。
我们用特定的值初始化这些变量。
我们使用带有struct文字的短变量声明操作符来初始化。
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
下面是我们如何谈论结构:
——用户定义类型
——我们声明类型。
类型有字段。
——类型也可以有“标签”
我们还没有看到这个。
——类型具有基础类型。
在这种情况下，底层类型是struct。
——我们声明类型的变量。
我们初始化这些变量。
——用特定的值进行初始化。
——或者，初始化为零值。
结构体是一种复合类型。
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
比尔·肯尼迪:
Go允许我们声明自己的类型。
结构类型通过组合一组固定的唯一字段来声明。
结构中的每个字段都以已知类型声明。
这可能是内置类型或其他用户定义的类型。
一旦我们声明了类型，我们就可以从类型中创建值。
当我们声明变量时，变量表示的值总是被初始化。
值可以用一个特定的值初始化，也可以初始化为零值。
对于数值类型，零值为0;对于弦，它是空的;
对于布尔人来说，这是错误的。
在结构体的情况下，零值适用于结构中的所有不同字段。
在创建变量并将其初始化为零值时，使用关键字var是惯用法。
保留关键字var的使用，以表明变量被设置为其零值。
如果变量的初始值为0，
然后使用具有结构文字的短变量声明操作符。
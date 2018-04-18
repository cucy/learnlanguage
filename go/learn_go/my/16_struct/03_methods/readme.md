(p person) is the "receiver"
it is another parameter
not idiomatic to use "this" or "self"


"Not many people know this, but method notation, i.e. v.Method() is actually syntactic sugar and Go also understands the de-sugared version of it: (T).Method(v). You can see an example here. Naming the receiver like any other parameter reflects that it is, in fact, just another parameter quite well.
This also implies that the receiver-argument inside a method may be nil. This is not the case with this in e.g. Java."
SOURCE:
https://www.reddit.com/r/golang/comments/3qoo36/question_why_is_self_or_this_not_considered_a/?utm_source=golangweekly&utm_medium=email


------------------------



(p人)是“接受者”
这是另一个参数
不习惯使用"this"或"self"


“没有多少人知道这一点，但是方法表示法，也就是v. method()实际上是语法糖，Go也理解了它的去糖版本:(T). method (v)。你可以在这里看到一个例子。像其他参数一样命名接收方反映了它实际上只是另一个参数。
这也意味着方法内的接收参数可能为nil。这在Java中不是这样的。
来源:
https://www.reddit.com/r/golang/comments/3qoo36/question_why_is_self_or_this_not_considered_a/?utm_source=golangweekly&utm_medium=email
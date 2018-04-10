"""
Chaining coroutines
In certain situations, you may wish to chain the calling of your coroutines together in order to achieve maximum performance within your Python systems.

The official documentation again has an excellent code sample that demonstrates this concept of chaining very well. Within this code, we have two distinct coroutines denoted by async def. The compute coroutine returns the summation of x + y after having performed a blocking sleep for 1 second.

Let's assume, however, that we want to rely on the result of a second coroutine within our first coroutine as follows:


链协同程序
在某些情况下，您可能希望将您的coroutines的调用链接在一起，以便在您的Python系统中实现最大的性能。

官方文档再次有一个优秀的代码示例，它演示了链接的概念。在这段代码中，我们有两个不同的coroutines，由async def表示。计算的coroutine在执行阻塞睡眠1秒后返回x + y的总和。

但是，让我们假设，我们想要依赖于我们第一个coroutine中的第二个coroutine的结果如下:
"""

import asyncio


async def compute(x, y):
    print("Compute %s + %s ..." % (x, y))
    await asyncio.sleep(1.0)
    return x + y


async def print_sum(x, y):
    result = compute(x, y)
    print("%s + %s = %s" % (x, y, result))


loop = asyncio.get_event_loop()
loop.run_until_complete(print_sum(1, 2))
loop.close()


"""
This essentially states that the compute function within our print_sum() method call was never awaited and the program tried to carry on as if it had received the result.

In order to overcome this particular issue, we need to utilize the await keyword. This await keyword blocks the event loop from proceeding any further until the called coroutine returns its result. The main drawback of this, however, is that, in this particular example, we lose the benefits of asynchronicity and we are back to standard synchronous execution. It's up to you to determine where the use of await is necessary as it does give you very quick and easy deterministic execution but you take hits on performance.

The print_sum coroutine instantiates a result variable that is equal to whatever the compute coroutine returns:

这本质上说明了print_sum()方法调用中的计算函数从来没有被等待过，程序试图继续进行，就好像它已经收到了结果一样。

为了克服这个特殊的问题，我们需要利用等待关键字。这个等待关键字阻止事件循环，直到被调用的coroutine返回结果。然而，这其中的主要缺点是，在这个特殊的例子中，我们失去了异步的好处，我们又回到了标准的同步执行。由你来决定等待的使用是必要的，因为它确实给了你非常快速和容易的确定性的执行，但是你会对性能产生影响。

print_sum coroutine实例化一个结果变量，它等于任何计算相关的结果:

"""

import asyncio

async def compute(x, y):
    print("Compute %s + %s ..." % (x, y))
    await asyncio.sleep(1.0)
    return x + y

async def print_sum(x, y):
    result = await compute(x, y)
    print("%s + %s = %s" % (x, y, result))

loop = asyncio.get_event_loop()
loop.run_until_complete(print_sum(1, 2))
loop.close()



"""
如果我们尝试运行前面的程序，您应该看到我们的coroutines已经成功地链接在一起，
并且我们的print_result coroutine现在正在等待我们的计算函数的结果:
$ python3.6 06_chainCoroutine.py
Compute 1 + 2 ...
1 + 2 = 3

Source: https://docs.python.org/3/library/asyncio-task.html
"""





"""
Futures
We've already been exposed to one kind of future object when we were looking in depth at executors and pools in Chapter 7, Executors and Pools. The asyncio module, however, provides a slightly different implementation with which we can play in our Python programs.

The asyncio.futures.Future object implementation of futures is almost identical to the concurrent.futures.Future object in terms of the underlying concepts. They are created with the intention that they will eventually be given a result some time in the future. They also feature all of the same methods that your standard future objects have.

There are, however, three distinct changes that have been outlined in the official documentation, which can be found at https://docs.python.org/3/library/asyncio-eventloop.html#futures. These changes are worth noting to avoid frustration in the future:

The result () and exception() function do not take a timeout parameter and raise an exception when the future isn't done yet
Callbacks registered with the add_done_callback() format are always called through the call_soon_threadsafe()event loops
The future class is not compatible with the wait and as_completed functions in the concurrent.futures package


期货
我们已经接触过一种未来的对象当我们在第7章，执行器和池中深入研究执行器和池时。然而，asyncio模块提供了一个稍微不同的实现，我们可以在Python程序中使用它。

asyncio.futures。期货的未来目标的实现几乎和当前的期货一样。关于潜在概念的未来目标。他们被创造出来的目的是，他们最终会在将来某个时候得到一个结果。它们还包含了您的标准未来对象所拥有的所有相同的方法。

不过，官方文档中已经列出了三种不同的变化，它们可以在https://docs.python.org/3/library/asyncio-eventloop.html#futures中找到。这些变化值得注意，以避免未来的挫折:

result()和exception()函数不接受超时参数，并且在将来还未完成时引发异常。
使用add_done_callback()格式注册的回调总是通过call_soon_threadsafe()事件循环调用。
将来的类与concurrent中的wait和as_completed函数不兼容。期货包

"""

"""
In this next example, we will look at how we can wrap a coroutine that we defined as myCoroutine in a future object and then handle it as such within a main coroutine that we'll define.

In this example, it's important to note the use of await before our call to asyncio.ensure_future. The ensure_future() method is the method that both schedules our coroutine for execution while also wrapping it in a future, so we have to ensure that this is completed before we try to access result():


在下一个示例中，我们将研究如何包装一个coroutine，我们将它定义为未来对象中的myCoroutine，然后在我们将定义的main coroutine中处理它。

在本例中，重要的是在我们调用asyncio.ensure_future之前注意等待的使用。ensure_future()方法是将我们的coroutine调度为执行的方法，同时也将其包装在将来，因此我们必须确保在尝试访问result()之前完成此操作:

"""

import asyncio


async def myCoroutine(future):
    await asyncio.sleep(1)
    future.set_result("My Future Has Completed")


async def main():
    future = asyncio.Future()
    await asyncio.ensure_future(myCoroutine(future))
    print(future.result())


loop = asyncio.get_event_loop()
try:
    loop.run_until_complete(main())
finally:
    loop.close()


"""
在执行前面的程序时，您应该看到我们的coroutine已经成功地转换为一个future对象，
并且我们可以使用我们通常希望使用的相同的.result()方法来访问它:
"""

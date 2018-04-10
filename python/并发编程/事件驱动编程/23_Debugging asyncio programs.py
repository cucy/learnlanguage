"""
Debugging asyncio programs
Thankfully, when it comes to debugging asyncio-based applications, we have a couple of options to consider. The writers of the asyncio module have very kindly provided a debug mode, which is quite powerful and can really aid us in our debugging adventures without the overhead of modifying the system's code base too dramatically.

"""

"""
Turning on this debug mode within your asyncio-based programs is relatively simply and requires just a call to this function:
在基于异步的程序中打开这个调试模式相对简单，只需调用这个函数:
loop.set_debug(True)


让我们来看看一个完整的示例，以及它与您的标准日志记录的不同之处。
在本例中，我们将创建一个非常简单的事件循环，并向事件循环提交一些简单的任务:

"""

import asyncio
import logging
import time

logging.basicConfig(level=logging.DEBUG)


async def myWorker():
    logging.info("My Worker Coroutine Hit")
    time.sleep(1)


async def main():
    logging.debug("My Main Function Hit")
    await asyncio.wait([myWorker()])


loop = asyncio.get_event_loop()
loop.set_debug(True)
try:
    loop.run_until_complete(main())
finally:
    loop.close()


"""
您应该注意到日志中没有包含的额外日志语句。这些额外的日志语句为我们提供了关于事件循环的更细粒度的概念。

在这个简单的示例中，调试模式能够捕捉到一个事实，即我们的一个coroutines花费了超过100毫秒的时间来执行，而当我们的事件循环最终关闭时。

您应该注意到，这种asyncio调试模式非常适合做一些事情，比如确定什么是coroutines永远不会被放弃，从而扼杀您的程序的性能。其他可以执行的检查如下:

call_soon()和call_at()方法在从错误的线程调用时引发异常。
记录选择器的执行时间。
当传输和事件循环未显式关闭时，资源警告警告就会发出。
这只是开始。总的来说，这是一种非常强大的工具，它可以让您选择除了典型的Pdb之外的其他选项，以便更深入地了解您的程序。
我希望您可以在https://docs.python.org/3/library/asyncio-dev.html#asyncio-debug-mode下查看关于asyncio的调试模式的官方文档，
因为它提供了一些更深入的示例，您可以在自己的空闲时间仔细阅读。
"""
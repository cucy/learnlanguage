"""
任务
asyncio内的任务负责在事件循环中执行协同程序。这些任务一次只能在一个事件循环中运行，为了实现并行执行，您必须在多个线程上运行多个事件循环。

我喜欢考虑在asyncio内部的任务，类似于我们如何看待与执行器或池一起使用时的任务，就像我们在前几章中演示的那样。

在本节中，我们将讨论一些关键函数，以便在基于异步的程序中处理任务。


在本例中，我们将研究如何定义一个生成器函数，该函数将为我们的事件循环生成5个不同的coroutines，以便调度和执行。为了调度这些协同程序，我们将使用ensure_future()方法，您将在本章进一步了解更多细节:
"""

import asyncio
import time


@asyncio.coroutine
def myTask(n):
    time.sleep(1)
    print("Processing {}".format(n))


@asyncio.coroutine
def myGenerator():
    for i in range(5):
        asyncio.ensure_future(myTask(i))
    print("Completed Tasks")
    yield from asyncio.sleep(2)


def main():
    loop = asyncio.get_event_loop()
    loop.run_until_complete(myGenerator())
    loop.close()


if __name__ == '__main__':
    main()

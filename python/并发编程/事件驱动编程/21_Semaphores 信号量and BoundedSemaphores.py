"""
Semaphores and BoundedSemaphores
信号量和有界信号量
asyncio模块为我们提供了它自己的信号量原语的实现，以及在线程模块中也可以找到的BoundedSemaphore原语。

这给我们提供了与我们在第4章中看到的线程同步的相同功能。它是一个计数器，随着每次调用的获得和增加而递减。同样，这允许我们通过访问给定的资源来控制coroutines的数量，并且可以帮助确保资源匮乏等问题不会成为我们程序中的问题。

这些可以实例化如下:

import asyncio
...
mySemaphore = asyncio.Semaphore(value=4, *, loop=None)
...
boundedSemaphore = asyncio.BoundedSemaphore(value=4, *, loop=None)
"""
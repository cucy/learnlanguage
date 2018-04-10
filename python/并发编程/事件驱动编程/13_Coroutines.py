"""
Coroutines in asyncio are similar to the standard Thread object that you'd find within the threading module. By utilizing coroutines within our asyncio-based application, we are essentially enabling ourselves to write asynchronous programs with the main exception that they run in a single-threaded context.

They are quite possibly the most important part of the asyncio module as they are typically where the magic happens within your event-based programs. If you look at any major asyncio-based program, you should notice a heavy utilization of these coroutine objects.

There are a couple of different ways we can implement our own coroutines, the first of which is to implement an async def function, which is a feature added to Python 3.5 and is definitely the method I recommend the most. If we were to implement a coroutine using this method, it would look something like this:

asyncio中的Coroutines类似于您在线程模块中找到的标准线程对象。通过在我们的基于异步的应用程序中使用coroutines，我们本质上允许自己编写异步程序，主要例外是它们在单线程上下文中运行。

它们很可能是asyncio模块中最重要的部分，因为它们通常是在基于事件的程序中发生的奇迹。如果您查看任何主要的基于异步的程序，您应该注意到这些coroutine对象的大量使用。

有几种不同的方法可以实现我们自己的coroutines，第一个是实现一个async def函数，这是Python 3.5中添加的一个特性，这绝对是我推荐的方法。如果我们用这个方法来实现一个coroutine，它看起来是这样的:
"""

import asyncio


async def myCoroutine():
    print("My Coroutine")


def main():
    loop = asyncio.get_event_loop()
    loop.run_until_complete(myCoroutine())
    loop.close()


if __name__ == '__main__':
    main()



# The second method is to utilize generators in conjunction with the @asyncio.coroutine decorator:

import asyncio

@asyncio.coroutine
def myCoroutine():
   print("My Coroutine")

def main():
   loop = asyncio.get_event_loop()
   loop.run_until_complete(myCoroutine())
   loop.close()

if __name__ == '__main__':
   main()


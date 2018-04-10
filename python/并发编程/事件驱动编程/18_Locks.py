"""
与您将在线程模块中找到的标准锁实现相比，异步模块中的锁在实用功能方面几乎是相同的。

asyncio模块的实现使我们可以在异步的程序的关键部分中设置锁，以确保没有其他的协同程序同时执行相同的关键部分。

为了从asyncio模块中实例化一个锁，我们必须这样做:

Import asyncio
myLock = asyncio.Lock()


Let's take a look at a fully fledged example of this. We'll begin by defining an async coroutine called myWorker, which will take in our lock as our parameter. Within this myWorker coroutine, we will attempt to acquire the lock using the with keyword. Once we've attained this lock, we then execute our critical status, which in this instance happens to be a simple print statement.

Within our main coroutine, we then instantiate the lock that we'll be passing to our coroutine function. We then await the execution of two instances of our myWorker coroutine before completing:


让我们来看看一个完全成熟的例子。我们将首先定义一个名为myWorker的async coroutine，它将把我们的锁作为参数。在这个myWorker coroutine中，我们将尝试使用with关键字来获取锁。一旦我们获得了这个锁，我们就会执行我们的临界状态，在这个实例中，它恰好是一个简单的print语句。

在我们的主coroutine中，然后实例化我们将传递给我们的coroutine函数的锁。然后我们等待我们的myWorker coroutine的两个实例在完成之前执行:

"""

import asyncio
import time


async def myWorker(lock):
    with await lock:
        print(lock)
        print("myWorker has attained lock, modifying variable")
        time.sleep(2)
    print(lock)
    print("myWorker has release the lock")


async def main(loop):
    lock = asyncio.Lock()
    await asyncio.wait([myWorker(lock),
                        myWorker(lock)])


loop = asyncio.get_event_loop()
try:
    loop.run_until_complete(main(loop))
finally:
    loop.close()


"""
<asyncio.locks.Lock object at 0x0000000002DFBC88 [locked]>
myWorker has attained lock, modifying variable
<asyncio.locks.Lock object at 0x0000000002DFBC88 [unlocked]>
myWorker has release the lock
<asyncio.locks.Lock object at 0x0000000002DFBC88 [locked]>
myWorker has attained lock, modifying variable
<asyncio.locks.Lock object at 0x0000000002DFBC88 [unlocked]>
myWorker has release the lock


如果我们要执行这个程序，您应该看到每个工人依次获得锁并执行它必须执行的操作，然后释放该锁，
并有效地允许我们的第二个coroutine接管并执行它自己的关键代码段:
"""

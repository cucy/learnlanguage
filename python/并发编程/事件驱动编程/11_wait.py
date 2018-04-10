"""
等待函数简单地阻塞了我们的程序，直到所有的期货或协同程序传递到该函数的第一个参数，成功完成:


在执行这个程序时，您应该看到我们已经成功地创建了我们的四个coroutines，并且在我们的程序终止之前它们都已经成功地完成了:
"""

import asyncio


async def myCoroutine(i):
    print("My Coroutine {}".format(i))


loop = asyncio.get_event_loop()
try:
    tasks = []
    for i in range(4):
        tasks.append(myCoroutine(i))
    loop.run_until_complete(asyncio.wait(tasks))
finally:
    loop.close()

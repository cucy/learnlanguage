"""
all_tasks方法返回给定事件循环的一组任务。如果没有传入事件循环，则默认只显示当前事件循环的所有任务:

在执行前面的程序时，您应该会看到一组输出到控制台的三种代码。
这些表示为pending，因为它们还没有被安排在我们当前的事件循环中运行:
"""

import asyncio


async def myCoroutine():
    print("My Coroutine")


async def main():
    await asyncio.sleep(1)


loop = asyncio.get_event_loop()
try:
    loop.create_task(myCoroutine())
    loop.create_task(myCoroutine())
    loop.create_task(myCoroutine())

    pending = asyncio.Task.all_tasks()
    print(pending)
    loop.run_until_complete(main())
finally:
    loop.close()

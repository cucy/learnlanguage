"""
测量当前任务的执行情况在许多情况下都是有用的。如果需要，您可以有效地遍历事件循环中执行的当前任务列表，并尝试在您希望的情况下取消它们。

在本例中，我们将使用create_task函数安排三个不同的任务，并将myCoroutine函数作为其输入:

在执行过程中，您应该会看到coroutines都已被成功执行，并且在我们的主coroutine被执行的时候，唯一挂起的coroutine是当前正在执行的一个:

"""

import asyncio


async def myCoroutine():
    print("My Coroutine")


async def main():
    current = asyncio.Task.current_task()
    print(current)


loop = asyncio.get_event_loop()
try:
    loop.create_task(myCoroutine())
    loop.create_task(myCoroutine())
    loop.create_task(myCoroutine())
    loop.run_until_complete(main())
finally:
    loop.close()

"""
取消功能允许我们要求取消期货或协程:


"""

import asyncio

async def myCoroutine():
    print("My Coroutine")

async def main():
    current = asyncio.Task.current_task()
    print(current)

loop = asyncio.get_event_loop()
try:
    task1 = loop.create_task(myCoroutine())
    task2 = loop.create_task(myCoroutine())
    task3 = loop.create_task(myCoroutine())
    task3.cancel()
    loop.run_until_complete(main())
finally:
    loop.close()

"""
My Coroutine
My Coroutine
<Task pending coro=<main() running at C:/Users/zhourudong/PycharmProjects/learn/事件驱动编程/9_cancel.py:14> cb=[_run_until_complete_cb() at C:\3.6\lib\asyncio\base_events.py:176]>


在执行前面的程序时，您应该看到task1和task2都已成功执行。我们计划的第三个任务，由于取消了我们的调用，实际上从来没有执行过。现在，这只是一个简单的例子，说明我们如何取消一个任务，我们以这样的方式做了，我们几乎可以保证我们的第三个任务被取消了。然而，在野外，不能保证取消功能肯定会取消您的待定任务:
"""
"""
run_until_complete方法是我们在本章中使用过的方法，它允许我们在结束自己之前给事件循环一个特定的工作量:

前面的代码将启动一个事件循环，然后在完成该函数之前执行myWork() async函数。
"""

import asyncio
import time


async def myWork():
    print("Starting Work")
    time.sleep(5)
    print("Ending Work")


def main():
    loop = asyncio.get_event_loop()
    loop.run_until_complete(myWork())
    loop.close()


if __name__ == '__main__':
    main()


"""
Starting Work
Ending Work
"""
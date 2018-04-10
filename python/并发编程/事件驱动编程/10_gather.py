"""
The gather function is a somewhat complicated beast. It takes in a set of coroutines or futures and then returns a future object that aggregates the results from the inputted set:

集合函数是一个有点复杂的野兽。它接受一组相关的或未来的，然后返回一个未来的对象，它聚集来自输入的集合的结果:
"""

import asyncio

async def myCoroutine(i):
    print("My Coroutine {}".format(i))

loop = asyncio.get_event_loop()
try:
    loop.run_until_complete(asyncio.gather(myCoroutine(1), myCoroutine(2)))
finally:
    loop.close()
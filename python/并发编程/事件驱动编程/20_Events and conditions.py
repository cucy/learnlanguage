"""
Events, as we have covered in Chapter 4, Synchronization between Threads, allow us to block multiple consumers from progressing from a set point until a flag has been set. We can instantiate them as follows:

正如我们在第4章中所讨论的，事件之间的同步，允许我们阻止多个消费者从一个设置点前进，直到一个标志被设置。我们可以实例化它们如下:
myEvent = asyncio.Event()


Conditions again allow us to block tasks until a point where they are notified that they can continue by another coroutine. We can instantiate this as follows:

条件再次允许我们阻塞任务，直到它们被通知可以继续使用另一个coroutine。我们可以将其实例化如下:

myCondition = asyncio.Condition()
"""
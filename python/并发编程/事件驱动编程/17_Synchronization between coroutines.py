"""

Synchronization between coroutines
The coroutines in asyncio run in a non-deterministic manner, and as such, there is still the possibility of our code running into race conditions regardless of the fact that it's only ever running on a single thread.

Due to the fact we can still be hit by dreaded race conditions, it's important to know how we can use the previously covered synchronization primitives in conjunction with the asyncio module.

In this section, we'll take a look at locks, events, conditions, and queues, as well as look at small examples of how these work within our event-driven Python programs. In the interest of brevity, I'm only going to give full code examples for both locks and queues as I feel these are going to be the two most commonly used concepts when it comes to synchronization between your asyncio coroutines.


asyncio中的coroutines以一种不确定的方式运行，因此，我们的代码仍然有可能运行到竞争环境中，而不考虑它只在单个线程上运行的事实。

由于我们仍然会受到可怕的竞争环境的影响，所以很重要的一点是，我们应该知道如何将以前覆盖的同步原语与asyncio模块一起使用。

在本节中，我们将查看锁、事件、条件和队列，并查看在事件驱动的Python程序中这些工作的示例。为了简洁起见，我只会给出两个锁和队列的完整代码示例，因为我觉得这将是两个最常用的概念，当它涉及到您的asyncio coroutines之间的同步时。
"""
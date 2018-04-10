"""
The asyncio.Queue implementation gives us another near identical implementation to that given to us in the standard queue module already present in Python.

In this example, we'll create both a producer and a consumer. The producer will produce articleIds between 1 and 5, while the consumer will try to read all of these articles. For both the producer and consumer, we'll define an appropriately named coroutine. Within our newsProducer coroutine, we'll perform the put command once every second that will put an ID onto our passed-in asyncio.Queue.

Within our newsConsumer function, we'll constantly try to perform a get request in order to retrieve an article ID from our shared queue. We'll then print out that we've consumed the said article ID and proceed to try and consume the next article:

asyncio。队列实现向我们提供了另一个几乎相同的实现，这是在Python中已经存在的标准队列模块中给出的。

在本例中，我们将创建一个生产者和一个消费者。生产者将在1到5之间生成articleid，而消费者将尝试读取所有这些文章。对于生产者和消费者，
我们将定义一个合适的coroutine。在我们的新闻生成器coroutine中，我们将每秒钟执行一次put命令，它将把ID放到我们的传入异步队列中。

在我们的newsConsumer函数中，我们将不断地尝试执行get请求，以便从共享队列检索文章ID。
然后，我们将打印出我们已经使用了上述文章ID并继续尝试并使用下一篇文章:

"""
import asyncio
import random
import time


@asyncio.coroutine
def newsProducer(myQueue):
    while True:
        yield from myQueue.put(random.randint(1, 5))
        yield from asyncio.sleep(1)


@asyncio.coroutine
def newsConsumer(myQueue):
    while True:
        articleId = yield from myQueue.get()
        print("News Reader Consumed News Article {}", articleId)


myQueue = asyncio.Queue()

loop = asyncio.get_event_loop()

loop.create_task(newsProducer(myQueue))
loop.create_task(newsConsumer(myQueue))
try:
    loop.run_forever()
finally:
    loop.close()


"""
When we execute the preceding program, you should see a constant stream of print statements coming from our newsConsumer coroutine. This demonstrates that we've successfully utilized our asyncio.Queue object for both our producer and consumer and have now achieved the communication that we desired:


当我们执行前面的程序时，您应该会看到来自我们的newsConsumer coroutine的连续的打印语句流。
这表明我们已经成功地利用了我们的asyncio。我们的生产者和消费者的队列对象现在已经实现了我们所期望的通信:
"""
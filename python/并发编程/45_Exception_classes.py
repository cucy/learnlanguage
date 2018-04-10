"""
在前一章中，我们讨论了如何通过使用队列原语来处理Python中的正常子线程中的异常，以便将异常从子线程传递到另一个线程。然而，这在某种程度上是一种过时的做法，而且，谢天谢地，对于threadpoolexecutor，我们不再需要担心这个问题，因为它已经为我们处理过了。

就像我们从我们的未来对象中检索结果一样，我们也可以返回异常。
"""

from concurrent.futures import ThreadPoolExecutor
import concurrent.futures
import threading
import random


def isEven(n):
    print("Checking if {} is even".format(n))
    if type(n) != int:
        raise Exception("Value entered is not an integer")
    if n % 2 == 0:
        print("{} is even".format(n))
        return True
    else:
        print("{} is odd".format(n))
        return False


def main():
    with ThreadPoolExecutor(max_workers=4) as executor:
        task1 = executor.submit(isEven, (2))
        task2 = executor.submit(isEven, (3))
        task3 = executor.submit(isEven, ('t'))

    for future in concurrent.futures.as_completed([task1, task2, task3]):
        print("Result of Task: {}".format(future.result()))


if __name__ == '__main__':
    main()

import time
import random
from concurrent.futures import ThreadPoolExecutor


def someTask(n):
    print("Executing Task {}".format(n))
    time.sleep(n)
    print("Task {} Finished Executing".format(n))


def main():
    with ThreadPoolExecutor(max_workers=2) as executor:
        task1 = executor.submit(someTask, (1))
        task2 = executor.submit(someTask, (2))
        executor.shutdown(wait=True)
        task3 = executor.submit(someTask, (3))
        task4 = executor.submit(someTask, (4))


if __name__ == '__main__':
    main()


"""
例子
在本例中，我们将演示运行执行器的关闭。我们首先定义一个函数，它本质上是“工作”的n秒。我们将提交一些任务，
然后调用我们的executor上的shutdown方法。在此之后，我们将尝试向执行人提交更多的任务:

"""
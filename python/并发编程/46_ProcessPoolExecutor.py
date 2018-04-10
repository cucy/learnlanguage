"""
  下面的示例提供了一个非常简单的完整示例，说明如何实例化自己的ProcessPoolExecutor，并将几个任务提交到这个池中。应该注意的是，我们的任务函数并不是计算成本很高的，所以我们可能看不到使用多个进程的全部好处，实际上，它可能比典型的单线程进程要慢得多。

我们将使用os模块来找到我们在池中执行的每个任务的当前PID:

"""
from concurrent.futures import ProcessPoolExecutor
import os


def task():
    print("Executing our Task on Process {}".format(os.getpid()))


def main():
    executor = ProcessPoolExecutor(max_workers=3)
    task1 = executor.submit(task)
    task2 = executor.submit(task)


if __name__ == '__main__':
    main()

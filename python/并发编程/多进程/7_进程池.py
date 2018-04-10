"""
concurrent.futures之间的区别。ProcessPoolExecutor和游泳池
我们concurrent.futures覆盖。在第7章，执行器和池中，ProcessPoolExecutors，那么，一个过程池的另一个实现的需要是什么?

多处理。过程池的池实现利用几乎相同的实现来提供并行处理功能。然而，并发的目的。futures模块是在创建流程池时提供一个更简单的接口。这个简单的界面很容易让程序员立即开始使用线程和进程池。然而，由于这种抽象的复杂性，我们在某些特定场景中可能需要的更细粒度的控制中失去了一些。

由于ThreadPoolExecutor和ProcessPoolExecutor是同一个抽象类的子类，所以与它们一起工作并记住它们的继承方法也要容易得多。

当涉及到Python 2和Python 3的可用性时，多处理模块胜过并发。期货，因为它是在2.6版本的语言中引入的，并且您不需要使用反向移植的版本。

一般来说，我建议并发。多进程的期货模块。池模块，因为它将满足您的需求，如果不是全部的话，大部分时间。然而，你确实需要知道，在那个决定命运的日子到来的时候，你是否会遇到同样的局限性。期货需要更多的控制。
"""

from multiprocessing import Pool


def task(n):
    print(n)


def main():
    with Pool(4) as p:
        print(p.map(task, [2, 3, 4]))


if __name__ == '__main__':
    main()

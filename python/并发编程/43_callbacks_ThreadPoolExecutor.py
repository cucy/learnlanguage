"""
设置回调函数
在第9章，事件驱动编程和第10章反应性编程中，我们将会详细介绍回调，当我们看事件驱动编程和反应性编程时。可视化回调的最好方法是想象你要求某人做某事，需要相当长的时间。通常情况下，当那个人完成任务时，你不会坐以待毙;你会去做其他的事情和你的时间。

相反，你会让他们在完成任务后给你回电话。在编程中，我们适当地将其称为回调，它们是一个非常强大的概念，我们可以与threadpoolexecutor一起使用。

在将任务提交给我们的ThreadPoolExecutor时，我们可以使用add_done_callback函数为上述函数指定一个回调函数:

ThreadPoolExecutor(max_workers = 3)作为遗嘱执行人:
未来=执行人。(2)提交(任务)
future.add_done_callback(taskDone)
这将确保在任务完成时调用taskDone函数。通过这个简单的函数调用，我们省去了观看我们提交自己的每一个任务的麻烦，然后，然后，在完成了一些其他的函数之后，我们就开始了。这些回调是非常强大的方式，他们为我们处理每一件事，我们几乎没有工作。
"""

"""
让我们看一看使用这个回调功能的完整样例代码。我们将从定义两个函数开始:任务函数，它将打印出我们传入它的任何东西，
以及taskDone函数，这将是我们的回调函数。

在这个taskDone函数中，我们首先检查我们的future对象是否已经被取消，或者是否已经完成。然后，我们将适当的输出输出到控制台。

在此基础上，我们定义了main函数，它简单地创建了ThreadPoolExecutor并将单个任务提交给它，同时还设置了它的回调:
"""

from concurrent.futures import ThreadPoolExecutor


def task(n):
    print("Processing {}".format(n))


def taskDone(fn):
    if fn.cancelled():
        "取消"
        print("Our {} Future has been cancelled".format(fn.arg))
    elif fn.done():
        print("Our Task has completed")


def main():
    print("Starting ThreadPoolExecutor")
    with ThreadPoolExecutor(max_workers=3) as executor:
        future = executor.submit(task, (2))
        future.add_done_callback(taskDone)

    print("All tasks complete")


if __name__ == '__main__':
    main()

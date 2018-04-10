import os


def child():
    print("We are in the child process with PID= %d" % os.getpid())


def parent():
    print("We are in the parent process with PID= %d" % os.getpid())
    newRef = os.fork()
    if newRef == 0:
        child()
    else:
        print("We are in the parent process and our child process has PID= %d" % newRef)


parent()

"""
We are in the parent process with PID= 8433
We are in the parent process and our child process has PID= 8477
We are in the child process with PID= 8477


把它分解
在前面的代码中，我们首先导入os Python模块。然后我们定义两个不同的函数，一个叫做子，一个叫父。子父进程只是打印进程标识符，否则称为PID。

在父函数中，我们在调用os.fork()方法来派生当前正在运行的进程之前，先打印出我们正在处理的进程的PID。这就创建了一个全新的过程，它接收到自己独特的PID。然后调用子函数，它输出当前的PID。您应该注意到，这个PID与我们脚本执行开始时打印出来的PID不同。

这个不同的PID代表了一个成功的forking和一个全新的过程。
"""
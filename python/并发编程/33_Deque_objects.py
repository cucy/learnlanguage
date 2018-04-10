"""
双端队列对象
Deques或double-end队列是另一种通信原语，我们可以积极地利用它来寻找线程安全的线程间通信。
它属于集合模块，它的功能非常类似于队列的功能，除了我们可以将元素插入到队列的两端。
"""

import collections

doubleEndedQueue = collections.deque('123456')

print("Dequeue: {}".format(doubleEndedQueue))

for item in doubleEndedQueue:
    print("Item {}".format(item))

print("Left Most Element: {}".format(doubleEndedQueue[0]))
print("Right Most Element: {}".format(doubleEndedQueue[-1]))

# 末尾插入
doubleEndedQueue.append('末尾')
print("Deque: {}".format(doubleEndedQueue))

# 前插入
doubleEndedQueue.appendleft('前边')
print("Deque: {}".format(doubleEndedQueue))

# 删除元素
# Removing Elements from our queue
rightPop = doubleEndedQueue.pop()
print(rightPop)
print("Deque: {}".format(doubleEndedQueue))

leftPop = doubleEndedQueue.popleft()
print(leftPop)
print("Deque: {}".format(doubleEndedQueue))

# 插入元素

print("插入元素前", "Deque: {}".format(doubleEndedQueue))

doubleEndedQueue.insert(5, "将数据插入到队列的第5个")

print("插入元素后", "Deque: {}".format(doubleEndedQueue))



# 反转队列  Source: http://www.transtutors.com/homework-help/accounting/inventory-valuation-lifo/

print("反转队列Deque: {}".format(doubleEndedQueue))

doubleEndedQueue.rotate(3)      #  向右拉4个位置

print("反转队列Deque: {}".format(doubleEndedQueue))

doubleEndedQueue.rotate(-2)

print("反转队列Deque {}".format(doubleEndedQueue))
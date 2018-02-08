[toc]
# 时间相关
```python

"""
pip install parsedatetime  arrow
"""

t0 = arrow.now()


```

# 命令行参数
```python


"""

pip installbegins
"""

import argparse


""" 
# v1



def main(a, b):
    return a + b


if __name__ == '__main__':
    parse = argparse.ArgumentParser(
        description="简单的两个数相加"
    )

    parse.add_argument("-a", help="第一个值", type=float, default=0)
    parse.add_argument("-b", help="第二个值", type=float, default=0)
    args = parse.parse_args()
    print(main(args.a, args.b))

'''
begins_lib.py -h

'''


"""

import begin


@begin.start(auto_convert=True)
def main(a: "第一个值" = 1.0, b: "第二个值" = 1.0):
    print(a + b)

```


# sched
```python



"""
pip install sched
"""

import sched
import time
from datetime import datetime, timedelta

"""  
# v1

scheduler = sched.scheduler(timefunc=time.time)


def saytime():
    print(time.ctime())
    scheduler.enter(1.5, priority=0, action=saytime)


saytime()
try:
    scheduler.run(blocking=True)
except KeyboardInterrupt:
    print("stoped 停止")


"""

# -------------     ##### v2


scheduler = sched.scheduler(timefunc=time.time)


def saytime():
    print(time.ctime(), flush=True)
    reschedule()


def reschedule():
    new_target = datetime.now().replace(
        second=0, microsecond=0
    )
    new_target += timedelta(minutes=1)  # 每分钟执行
    scheduler.enterabs(
        new_target.timestamp(), priority=0, action=saytime
    )


reschedule()
try:
    scheduler.run(blocking=True)
except KeyboardInterrupt:
    print("stoped. ")

```

# 日志
```python

import logging

logger = logging.getLogger()

logger.debug("debug message")
logger.info("info message")
logger.warning("warning message")
logger.error("error message")
logger.critical("critical message")

# ---------  -------------
import logging

logger = logging.getLogger()


def blah():
    return "blah"


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)

# ------- ---------
if __name__ == '__main__':
    from argparse import ArgumentParser

    parser = ArgumentParser(description="设置app日志等级")
    parser.add_argument("-ll", "--loglevel",
                        type=str,
                        choices=['DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL', ],
                        help="请按参数设置等级"

                        )

    """
    python logging_lib.py -ll  DEBUG

    """
```

```python

import colorlog

logger = colorlog.getLogger()
logger.setLevel(colorlog.colorlog.logging.DEBUG)

handler = colorlog.StreamHandler()
handler.setFormatter(colorlog.ColoredFormatter())

logger.addHandler(handler)

logger.debug("debug message")
logger.info("info message")
logger.warning("warning message")
logger.error("error message")
logger.critical("critical message")


```

# contextlib
```python

""" 
with open("test.txt", "r")  as f:
    data = f.read()
print(data)
"""

from time import perf_counter
from array import array
from contextlib import contextmanager


@contextmanager
def timing(label: str):
    t0 = perf_counter()
    yield lambda: (label, t1 - t0)
    t1 = perf_counter()


with timing("Arry test") as total:
    with timing("array creation innermul")  as inner:
        x = array("d", [0] * 100000)

    with timing("array creation outermul")  as outermul:
        x = array("d", [0] * 100000)

print("Total [%s]: %.6f s" % total())
print("  [%s]: %.6f s" % inner())
print("  [%s]: %.6f s" % outermul())

"""
Total [Arry test]: 0.014329 s
  [array creation innermul]: 0.007280 s
  [array creation outermul]: 0.006995 s
"""
#  -----------------  ---------------------

import threading


def work():
    """ 虽然简单 当参数过多后"""
    return sum(x for x in range(1000))


thread = threading.Thread(target=work)
thread.start()
thread.join()

#  -----------------  ---------------------
from concurrent.futures import ThreadPoolExecutor as Exexutor

urls = "baidu qq segmentfault alipay"


def fetch(url):
    from urllib import request, error
    try:
        print(url)
        data = request.urlopen(url).read()
        return "{}: length {} ".format(url, len(data))
    except error.HTTPError as e:
        return "{}: length {} ".format(url, e)


with Exexutor(max_workers=4) as exe:
    template = "http://www.{}.com"
    jobs = [exe.submit(
        fetch, template.format(u)) for u in urls.split()
    ]
    results = [job.result() for job in jobs]

print("\n".join(results))

```

# collections
```python
from string import ascii_lowercase

a = dict(zip(ascii_lowercase, range(10)))
print(a)
"""
{'j': 9, 'd': 3, 'e': 4, 'c': 2, 'g': 6, 'a': 0, 'i': 8, 'h': 7, 'f': 5, 'b': 1}

返回的并不是排序好的
"""

from collections import OrderedDict

res = OrderedDict(zip(ascii_lowercase, range(10)))
print(res)
"""
OrderedDict([('a', 0), ('b', 1), ('c', 2), ('d', 3), ('e', 4), ('f', 5), ('g', 6), ('h', 7), ('i', 8), ('j', 9)])
"""

# 定义简单的class
from collections import namedtuple

A = namedtuple("Person", "age sex color")
person = A(age=29, sex="M", color="yellow")

print(person.age)
print(person.sex)
print(person)
"""
29
M
Person(age=29, sex='M', color='yellow')

"""

#  -----------------  ---------------------
A = namedtuple("Person", "age sex color")


# 简写
def f():
    return A(29, "M", "yellow")


res = f()
print(res.color)  # yellow

#  -----------------  ---------------------

```

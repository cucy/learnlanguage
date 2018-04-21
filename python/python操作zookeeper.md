```python
pip kazoo
```

会话1
```python
In [1]: import kazoo

In [2]: from kazoo.client import KazooClient
In [2]: from kazoo.client import KazooClient

In [3]: zk = KazooClient(hosts='127.0.0.1:2181')

In [4]: zk.get_children("/")
In [5]: zk.start()

In [6]: zk.get_children("/")
Out[6]: ['zookeeper', 'cmdb']
In [6]:  zk.ensure_path("/cmdb/lock") #创建资源
In [7]: lock = zk.Lock("/cmdb/lock")

In [8]: lock.acquire() # 加锁
Out[8]: True

In [9]: lock.release() # 释放锁
Out[9]: True
```

会话2
```python

In [1]:  import kazoo

In [2]: rom kazoo.client import KazooClient
In [3]: from kazoo.client import KazooClient

In [4]: zk = KazooClient(hosts='127.0.0.1:2181')

In [5]: zk.start()

In [6]: lock = zk.Lock("/cmdb/lock")

In [7]: lock.acquire() # 会话2被堵塞等待其他锁释放
Out[7]: True
```

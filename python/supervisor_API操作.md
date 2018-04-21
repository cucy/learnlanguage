# 登录
server = xmlrpclib.Server('http://zrd:sh123!@192.168.1.60:9001/RPC2/')

## supervisor操作方法

### 获取版本
```python
In [28]: server.supervisor.getSupervisorVersion()
Out[28]: '3.3.2'
```
### 获取进程状态
```python
In [26]: server.supervisor.getState()
Out[26]: {'statecode': 1, 'statename': 'RUNNING'}
```
返回值:

| statecode | statename | Description |
| --------- | ---------- | ---------------------------------------- |
| 2 | FATAL | Supervisor has experienced a serious error. |
| 1 | RUNNING | Supervisor is working normally. |
| 0 | RESTARTING | Supervisor is in the process of restarting. |
| -1 | SHUTDOWN | Supervisor is in the process of shutting down. |


### 获取进程日志
```python
# readLog(offset, length)

In [34]: server.supervisor.readLog(1, 100)
Out[34]: '017-06-15 16:37:06,813 CRIT Supervisor running as root (no user in config file)\n2017-06-15 16:37:06,'
```
可选值:

| Offset | Length | Behavior of `readProcessLog` |
| ---------------- | -------- | ---------------------------------------- |
| Negative | Not Zero | Bad arguments. This will raise the fault `BAD_ARGUMENTS`. |
| Negative | Zero | This will return the tail of the log, or offset number of characters from the end of the log. For example, if `offset` = -4 and `length` = 0, then the last four characters will be returned from the end of the log. |
| Zero or Positive | Negative | Bad arguments. This will raise the fault `BAD_ARGUMENTS`. |
| Zero or Positive | Zero | All characters will be returned from the `offset` specified. |
| Zero or Positive | Positive | A number of characters length will be returned from the `offset`. |


### 清空日志
clearLog()
```python
In [35]: server.supervisor.clearLog()
Out[35]: True
```

### 关闭进程
```python
In [36]: server.supervisor.shutdown()
Out[36]: True
```

### 重启进程
```python
In [37]: server.supervisor.restart()
Out[37]: True
```


## supervisor管理其他进程操作方法

### 进程信息
```python
In [39]: server.supervisor.getProcessInfo('python_amon_gunicorn')
Out[39]:
{'description': 'pid 10709, uptime 0:02:41',
'exitstatus': 0,
'group': 'python_amon_gunicorn',
'logfile': '/tmp/python_amon_gunicorn.log',
'name': 'python_amon_gunicorn',
'now': 1497523086,
'pid': 10709,
'spawnerr': '',
'start': 1497522925,
'state': 20,
'statename': 'RUNNING',
'stderr_logfile': '',
'stdout_logfile': '/tmp/python_amon_gunicorn.log',
'stop': 0}
```

### 获取全部进程信息
```python
In [42]: server.supervisor.getAllProcessInfo()
Out[42]:
[{'description': 'pid 14036, uptime 0:01:29',
'exitstatus': 0,
'group': 'python_amon_gunicorn',
'logfile': '/tmp/python_amon_gunicorn.log',
'name': 'python_amon_gunicorn',
'now': 1497523775,
'pid': 14036,
'spawnerr': '',
'start': 1497523686,
'state': 20,
'statename': 'RUNNING',
'stderr_logfile': '',
'stdout_logfile': '/tmp/python_amon_gunicorn.log',
'stop': 0},
{'description': 'pid 14035, uptime 0:01:29',
'exitstatus': 0,
'group': 'zrd_mongod',
'logfile': '/tmp/zrd_mongod.log',
'name': 'zrd_mongod',
'now': 1497523775,
'pid': 14035,
'spawnerr': '',
'start': 1497523686,
'state': 20,
'statename': 'RUNNING',
'stderr_logfile': '',
'stdout_logfile': '/tmp/zrd_mongod.log',
'stop': 0}]
```

### 启动进程(所有)
```python
In [50]: server.supervisor.startAllProcesses()
Out[50]:
[{'description': 'OK',
'group': 'zrd_mongod',
'name': 'zrd_mongod',
'status': 80},
{'description': 'OK',
'group': 'python_amon_gunicorn',
'name': 'python_amon_gunicorn',
'status': 80}]
```
### 发送信号(停止单个进程)
```python
signalProcess(name, signal)
# 9 15
In [51]: server.supervisor.signalProcess('zrd_mongod', 15)
Out[51]: True

```
```python
# 用法同上
signalAllProcesses(signal)
```


### 打印日志
readProcessStdoutLog(name, offset, length)
```python
In [62]: print server.supervisor.readProcessStdoutLog('zrd_mongod', 0, 0)

```

### tail日志
tailProcessStdoutLog(name, offset, length)
```python
a = server.supervisor.tailProcessStdoutLog('zrd_mongod',0 , 10000)

In [73]: for i in a:
print i
```
### 清除日志
clearProcessLogs(name)
```python
In [74]: server.supervisor.clearProcessLogs('zrd_mongod')
Out[74]: True
```

### 清除所有日志
clearAllProcessLogs()
```python
In [75]: server.supervisor.clearAllProcessLogs()
Out[75]:
[{'description': 'OK',
'group': 'zrd_mongod',
'name': 'zrd_mongod',
'status': 80},
{'description': 'OK',
'group': 'python_amon_gunicorn',
'name': 'python_amon_gunicorn',
'status': 80}]
```



- 官方网址信息
http://supervisord.org/api.html

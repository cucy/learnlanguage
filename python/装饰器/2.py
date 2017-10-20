def log_cost_time(func):
    def wrapped(*args, **kwargs):
        import time
        begin = time.time()
        try:
            return func(*args,**kwargs)
        finally:
            print("函数{} 运行花费时间{}".format(
                func.__name__, time.time() - begin
            ))
    return wrapped

@log_cost_time
def complex_func(num):
    print(complex_func.__name__)
    ret = 0
    for i in range(num):
        ret += i * i
    return ret
print(complex_func(11111112))

"""
装饰器语法
@dec
def func():
    pass

等价于
func = dec(func)
"""

"""
副作用：
    被log_cost_time装饰的complex_func,
    print(complex_func.__name__)输出的是 wrapped
    这个是log_cost_time里面inner function(wrapped)的名字
    调用者希望输出的是complex_func，解决方法:
    
    1:
    functools.update_wrapper
    原型:  
     functools.update_wrapper(wrapper, wrapped[, assigned][, updated])
    第三个参数,将wrapped的值直接复制给wrapped,默认为:
      (__doc__, __name__, __module__)
    第四个参数, update, 默认为(__dict__)
    
"""
# 修改代码
print('\n')
import functools
def log_cost_time(func):
    @functools.wraps(func)
    def wrapped(*args, **kwargs):
        import time
        begin = time.time()
        try:
            return func(*args, **kwargs)
        finally:
            print("函数{} 运行用时 {}".format(
                func.__name__, time.time() -begin
            ))
    return wrapped
@log_cost_time
def complex_func(num):
    print(complex_func.__name__)
    ret = 0
    for i in range(num):
        ret += i * i
    return ret
print(complex_func(11111112))


###################################
# 带参数装饰器
###################################
print('\n')
def log_cost_time(stream):
    def inner_dec(func):
        def wrapped(*args, **kwargs):
            import time
            begin = time.time()
            try:
                return func(*args, **kwargs)
            finally:
                stream.write('函数 %s 费时 %s ' % (func.__name__, time.time() - begin))
        return wrapped
    return inner_dec

import sys
@log_cost_time(sys.stdout)
def complex_func(num):
    print(complex_func.__name__)
    ret = 0
    for i in range(num):
        ret += i * i
    return ret
print(complex_func(11111112))

"""
@dec(dec_args)
def func(*args, **kwargs):
    pass

<==> func = dec(dec_args)(*args, **kwargs)
"""



##############
# 类装饰器
##############
print('\n')
def Haha(clz):
    clz.__str__ = lambda s: 'Haha'
    return clz
@Haha
class Widget:
    """class Widget"""
w = Widget()
print(w)


##################
# 什么时候用装饰器
##################
"""
动态的为某个对象增加额外的责任
由于装饰器模式仅从外部改变组件，因此组件无需对它的装饰有任何了解；
也就是说，这些装饰对该组件是透明的。

1. 修改被装饰对象数或者行为
2. 出来被装饰函数对象执行的上下文, 比如设置环境变量,加log之类
3. 处理重复的逻辑,如有N个函数都肯跑出异常,但是我们不关心这些异常,
只要不向调用者传递异常就行，这时可以写一个catchall的decorator，
作用于可能跑出异常的函数
"""
print('\n')
import functools
def catchall(func):
    @functools.wraps(func)
    def wrapped(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except:
            pass
    return wrapped




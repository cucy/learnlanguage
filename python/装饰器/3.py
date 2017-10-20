def decorate(func):
    print("运行装饰器", func)
    def decorate_inner():
        print("运行 decorate_inner 函数")
        return func()
    return decorate_inner

@decorate
def func_1():
    print("运行 func_1")

print(func_1)
"""
运行装饰器 <function func_1 at 0x0000000000ABD400>
<function decorate.<locals>.decorate_inner at 0x0000000000ABD488>
"""

########################
# 如何使用被装饰函数中的参数
########################
print('\n')
def decorate(func):
    def decorate_inner(*args, **kwargs):
        print("参数类型: ",type(args), type(kwargs))
        print('args', args, 'kwargs', kwargs)
        return func(*args, **kwargs)
    return decorate_inner
@decorate
def func_1(*args, **kwargs):
    print(args, kwargs)

func_1(1,2,'3', parm1='a', parm2='b',parm3='c' )


#########################
#  叠放装饰器
# 装饰器的执行顺序
#########################
print('\n')
def outer(func):
    print("enter outer", func)
    def wrapper():
        print("运行 outer")
        func()
    return wrapper
def inner(func):
    print("enter inner", func)
    def wrapper():
        print("运行 inner")
        func()
    return wrapper
@outer
@inner
def main1():
    print("运行 main1 函数")
main1()

"""
enter inner <function main1 at 0x00000000010AD620>
enter outer <function inner.<locals>.wrapper at 0x00000000010AD6A8>
运行 outer
运行 inner
运行 main1 函数
"""

####################
# 消除装饰器的副作用
###################
print('\n')
from functools import wraps
def decorate(func):
    print("运行 decorate装饰器1111111",func )
    @wraps(func)
    def decorate_inner():
        print("运行 decorator_inner function", decorate_inner)
        return func()
    return decorate_inner
@decorate
def func_1():
    print("running func_1", func_1)
func_1()

# 一个装饰器是需要一个函数作为参数的函数
def my_shiny_new_decorator(a_function_to_decorate):
    """
    在装饰器内部动态定义一个函数: wrapper（原意:包装纸）
    # 这个函数将被包装在原始函数的四周
    # 因此就可以在原始函数之前和之后执行一些代码.
    """

    def the_wrapper_around_the_original_function():
        # 把想要在调用原始函数前运行的代码放这里
        print("Before the function runs")

        # 调用原始的函数(需要括号)
        a_function_to_decorate()

        # 把想要在调用原始函数后运行的代码放这里
        print("After the function runs")

        # 直到现在，"a_function_to_decorate"还没有执行过 (HAS NEVER BEEN EXECUTED).
        # 我们把刚刚创建的 wrapper 函数返回.
        # wrapper 函数包含了这个函数，还有一些需要提前后之后执行的代码，
        # 可以直接使用了（It's ready to use!）

    return the_wrapper_around_the_original_function
    # Now imagine you create a function you don't want to ever touch again.


def a_stand_alone_function():
    print("我是 a_stand_alone_function, 不要修改我!")


a_stand_alone_function()  # 我是 a_stand_alone_function, 不要修改我!
print('\n')
# 现在，你可以装饰一下来修改它的行为.
# 只要简单的把它传递给装饰器，后者能用任何你想要的代码动态的包装
# 而且返回一个可以直接使用的新函数:
a_stand_alone_function = my_shiny_new_decorator(a_stand_alone_function)
a_stand_alone_function()
print('\n')


# 使用装饰器语法
@my_shiny_new_decorator
def another_stand_alone_function():
    print ("别管我！"   )
another_stand_alone_function()
print('\n')



#######################################
# 装饰器的执行顺序
#######################################
def bread(func):
    """ 面包"""
    def wrapper():
        print("</ ''''''\>")
        func()
        print("<\______/>" )
    return wrapper

def ingredients(func):
    """ 成分 """
    def wrapper():
        print("西红柿")
        func()
        print("沙拉")
    return wrapper

@bread
@ingredients
def sanwich(food='--ham--'):
    """ 三明治 """
    print(food)
sanwich()
"""
</''''''\>
西红柿
--ham--
沙拉
<\______/>
"""


print('\n')
@ingredients
@bread
def sanwich(food='--ham--'):
    """ 三明治 """
    print(food)
sanwich()
"""
西红柿
</''''''\>
--ham--
<\______/>
沙拉
"""

#########################
# 装饰器传参
#########################
print('\n')
def a_decorator_passing_arguments(function_to_decorate):
    def a_wrapper_accepting_arguments(arg1, arg2):
        print("I got args! Lock", arg1, arg2)
        function_to_decorate(arg1,arg2)
    return a_wrapper_accepting_arguments
@a_decorator_passing_arguments
def print_full_name(first_name, last_name):
    print('My name is', first_name, last_name)

print_full_name('张', '三')

#################
# 带参数的装饰器
#################
print('\n')
def pre_str(pre=''):
    # old decorator
    def decorator(F):
        def new_F(a, b):
            print(  "{}input {} {}".format( pre, a, b ))
            return F(a, b)
        return new_F
    return decorator

# get square sum 求平方和
@pre_str('*_*')
def square_sum(a, b):
    return a ** 2 + b ** 2


# get square diff 求平方差
@pre_str('T_T')
def square_diff(a, b):
    return a ** 2 - b ** 2
print(square_sum(3, 4))
print(square_diff(3, 4))
"""
*_*input 3 4
25
T_Tinput 3 4
-7
"""

"""
square_sum = pre_str('*_*')(square_sum)
"""

#######
# 装饰方法
#######
print('\n')
def method_friendly_decorator(method_to_decorate):
    def wrapper(self, lie):
        lie = lie -3
        return method_to_decorate(self, lie)
    return wrapper
class Lucy:
    def __init__(self):
        self.age = 32

    @method_friendly_decorator
    def sayYourAge(self, lie):
        print("I am {} what did you think?".format(self.age, lie))

l = Lucy()
l.sayYourAge(-3)
""" I am 32 what did you think? """

#############33
# 通用的装饰器
#############
print('\n通用的装饰器\n')
def a_decorator_passing_arbitrary_arguments(function_to_decorate):
    def a_wrapper_accepting_arbitrary_arguments(*args, **kwargs):
        print("Do i have args? :")
        print(args)
        print(kwargs)
        function_to_decorate(*args, **kwargs)
    return a_wrapper_accepting_arbitrary_arguments
@a_decorator_passing_arbitrary_arguments
def function_with_no_argument():
    print('no argument here! 没有参数')
function_with_no_argument()
print('---------------')

@a_decorator_passing_arbitrary_arguments
def function_with_arguments(a,b, c):
    print(a, b, c)

function_with_arguments(1,2,3)
"""
Do i have args? :
(1, 2, 3)
{}
1 2 3
"""
print('---------------\n')
@a_decorator_passing_arbitrary_arguments
def function_with_named_aarguments(a, b,c, platypus="Why not?"):
    print("DO {}, {} and {} liek platypus? {}".format(
        a, b, c, platypus
    ))
function_with_named_aarguments('Bill', '李纳斯', '乔布斯', platypus="Indeed")
"""
Do i have args? :
('Bill', '李纳斯', '乔布斯')
{'platypus': 'Indeed'}
DO Bill, 李纳斯 and 乔布斯 liek platypus? Indeed
"""

print('---------------\n')
class Mary:
    def __init__(self):
        self.age = 31
    @a_decorator_passing_arbitrary_arguments
    def sayYourAge(self, lie=-3):
        print("I am {}, what did you think?".format(self.age + lie ))

m = Mary()
m.sayYourAge()
"""

Do i have args? :
(<__main__.Mary object at 0x00000000010FDDD8>,)
{}
I am 28, what did you think?

"""

############
# 类装饰器
############
print('\n------- 类装饰器 --------\n')
def decorator(aClass):
    class newClass:
        def __init__(self, age):
            self.total_display = 0
            self.wrapped = aClass(age)
        def display(self):
            self.total_display += 1
            print('total dispaly-> {}'.format(self.total_display))
            self.wrapped.display()
    return newClass
@decorator
class Bird:
    def __init__(self, age):
        self.age = age
    def display(self):
        print('My age is {}'.format(self.age))

eagLeLord = Bird(5)
for i in range(3):
    eagLeLord.display()
"""
total dispaly-> 1
My age is 5
total dispaly-> 2
My age is 5
total dispaly-> 3
My age is 5
"""
"""
在decorator中，我们返回了一个新类newClass。
在新类中，我们记录了原来类生成的对象（self.wrapped），
并附加了新的属性total_display，用于记录调用display的次数。
我们也同时更改了display方法。
通过修改，我们的Bird类可以显示调用display的次数了。
"""


##############
# 给装饰器传参（Passing arguments to the decorator）
#############
print('\n------- 给装饰器传参 --------\n')
def my_decorator(func):
    print("i am a ordinary function 普通函数")
    def wrapper():
        print("i am function returned by the decorator 我是一个函数由装饰器返回")
        func()
    return wrapper

def lazy_function():
    print('zzzzzzzzzz....')

decorated_function = my_decorator(lazy_function)
# i am a ordinary function 普通函数

@my_decorator
def lazy_function():
    print('zzzzzzzzzz....')
    # i am a ordinary function 普通函数


print('\n------- 给装饰器传参1 --------\n')
def decorator_maker_with_arguments(decorator_arg1, decorator_arg2):
    print ("I make decorators! And I accept arguments:{} {}".format(decorator_arg1, decorator_arg2   ))
    def my_decorator(func):
        # 在这里能传参数是一个来自闭包的馈赠.
        # 如果你对闭包感到不舒服，你可以直接忽略（you can assume it's ok）,
        # 或者看看这里: http://stackoverflow.com/questions/13857/can-you-explain-closures-as-they-relate-to-python
        print( "I am the decorator. Somehow you passed me arguments:", decorator_arg1, decorator_arg2   )

        # 不要把装饰器参数和函数参数搞混了！
        def wrapped(function_arg1, function_arg2) :
            print (
                "I am the wrapper around the decorated function.\n"  
                  "I can access all the variables\n"   
                  "\t- from the decorator: {0} {1}\n"   
                  "\t- from the function call: {2} {3}\n"   
                  "Then I can pass them to the decorated function"
                  .format(decorator_arg1, decorator_arg2,
                          function_arg1, function_arg2))
            return func(function_arg1, function_arg2)

        return wrapped

    return my_decorator

@decorator_maker_with_arguments("Leonard", "Sheldon")
def decorated_function_with_arguments(function_arg1, function_arg2):
    print("I am the decorated function and only knows about my arguments: {0}"   
           " {1}".format(function_arg1, function_arg2))

decorated_function_with_arguments("Rajesh", "Howard")

print('\n------- 给装饰器传参 区分装饰器参数和给函数传参数差别 2 --------\n')
c1 = '彼得'
c2 = '李磊'
@decorator_maker_with_arguments("彼得", c2)
def decorated_function_with_arguments(function_arg1, function_arg2):
    print("I am the decorated function and only knows about my arguments: {0}"   
           " {1}".format(function_arg1, function_arg2))

decorated_function_with_arguments("Rajesh", "Howard")


#############################
# 装饰器装饰一个装饰器
#############################
print('\n------- 装饰器装饰一个装饰器 --------\n')
def decorator_with_args(decorator_to_enhance):
    """
    This function is supposed to be used as a decorator. 这个函数应该用作装饰器。
    It must decorate an other function, 它必须装饰其他功能
    that is intended to be used as a decorator.
    Take a cup of coffee.  ，这是打算用作装饰。喝杯咖啡
    It will allow any decorator to accept an arbitrary number of arguments,
    saving you the headache to remember how to do that every time.
    它允许任何一个装饰师接受任意数量的参数，这样就可以节省每次记着如何做的头疼。
    """
    # We use the same trick we did to pass arguments   我们用同样的方法来传递参数。
    def decorator_maker(*args, **kwargs):
        # We create on the fly a decorator that accepts only a function
        # but keeps the passed arguments from the maker.
        """
        我们创建一个只接受一个函数的装饰器。
        但保持传递的参数从制造商。
        """

        def decorator_wrapper(func):
            # We return the result of the original decorator, which, after all,
            # IS JUST AN ORDINARY FUNCTION (which returns a function).
            # Only pitfall: the decorator must have this specific signature or it won't work:
            """
            我们返回原装饰器的结果，毕竟，
            只是一个普通函数（它返回一个函数）。
            只有陷阱：装饰器必须有这个特定的签名，否则它不会工作：
            """

            return decorator_to_enhance(func, *args, **kwargs)
        return decorator_wrapper
    return decorator_maker

@decorator_with_args
def decorated_decorator(func, *args, **kwargs):
    def wrapper(function_arg1, function_arg2):
        print("Decorator with", args, kwargs)
        return func(function_arg1, function_arg2)
    return wrapper
@decorated_decorator(42, 404, 1024)
def decorated_function(function_arg1, function_arg2):
    print('Hello {} {}'.format(function_arg1,function_arg2 ))
decorated_function("Universe and", "everything")
"""
Decorator with (42, 404, 1024) {}
Hello Universe and everything
"""

################################
# 装饰器最佳实践
################################
print('\n------- 副作用 --------\n')

def foo():
    print('foo')
print(foo.__name__) # foo

def bar(func):
    def wrapper():
        print('bar')
        return func()
    return wrapper
@bar
def foo():
    print('foo')
print(foo.__name__) # wrapper

# 解决方案
import functools
def bar(func):
    @functools.wraps(func)
    def wrapper():
        print('bar')
        return func()
    return wrapper
@bar
def foo():
    print('foo')
print(foo.__name__) # foo


###############################
# 装饰器如何才能有用
###############################
print('\n------- 装饰器如何才能有用 --------\n')

def benchmark(func):
    import time
    def wrapper(*args, **kwargs):
        t = time.clock()
        res = func(*args, **kwargs)
        print(func.__name__, time.clock() - t )
        return res
    return wrapper

def logging(func):
    def wrapper(*args, **kwargs):
        res = func(*kwargs, **kwargs)
        print(func.__name__, args, kwargs)
        return res
    return wrapper

def counter(func):
    def wrapper(*args, **kwargs):
        wrapper.count = wrapper.count + 1
        res = func(*args, **kwargs)
        print("{0} has been used: {1}x".format(func.__name__, wrapper.count))
        return res
    wrapper.count = 0
    return wrapper


@counter
@benchmark
@logging
def reverse_string(string):
    return str(reversed(string))


print(reverse_string("Able was I ere I saw Elba"))
print(reverse_string(
    "A man, a plan, a canoe, pasta, heros, rajahs, a coloratura, maps, snipe, percale, macaroni, a gag, a banana bag, a tan, a tag, a banana bag again (or a camel), a crepe, pins, Spam, a rut, a Rolo, cash, a jar, sore hats, a peon, a canal: Panama!"))

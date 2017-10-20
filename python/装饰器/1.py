
#########################3
# 函数也是对象
#########################
def message(word='hello'):
    return word.upper()

print(message())

# 函数也是对象
my_message = message
print(my_message())

###############
# 函数可以嵌套
###############


###############
# 函数用作参数返回
#############
def getWordType(kind='lower'):
    def capitalize(word='hello'):
        return word.capitalize()
    def lower(word='hello'):
        return word.lower()

    if kind == 'lower':
        return lower
    else:
        return capitalize
wordType = getWordType()
print(wordType)
print(wordType())

####################
# 函数作为参数传入
####################
def getName(name='leo'):
    return name
def foo(func):
    print('I will call the getName {} function later'.format(func.__name__))
    print(func())
foo(getName)



###########
# 装饰器
##########
"""
装饰器就是多函数进行再次包装,它在不改变原函数的前提下,增加函数的功能
可以在函数执行之前或者执行之后执行一段代码
"""
def my_new_decorator(a_function_to_decorate):
    def the_wrapper_around_the_original_function():
        print("Before the function runs, 调用 {} 前".format(a_function_to_decorate.__name__))
        a_function_to_decorate()
        print("After the functions runs,调用 {} 后".format(a_function_to_decorate.__name__))
    return the_wrapper_around_the_original_function

def a_stand_alone_funtion():
    print('I am a stand alone function,do not you dare modify me')

a_stand_alone_funtion()
a_stand_alone_funtion=my_new_decorator(a_stand_alone_funtion)
a_stand_alone_funtion()

# 换成装饰器写法
print('\n')
@my_new_decorator
def another_stand_alone_funtion():
    print("leave me alone")
another_stand_alone_funtion()  # another_stand_alone_funtion = my_new_decorator(another_stand_alone_funtion)

######################
# 为何使用装饰器
#####################
print('\n')
def makebole(func):
    """ 将字体变为粗体 """
    def wrapper():
        return "<b>{}</b>".format(func())
    return wrapper
def makeitalic(func):
    def wrapper():
        return '<i>{}</i>'.format(func())
    return wrapper

@makebole   #  2.再执行这个装饰器的内容
@makeitalic #  1.先执行这个装饰器的内容
def world():
    return "hello"
print(world())
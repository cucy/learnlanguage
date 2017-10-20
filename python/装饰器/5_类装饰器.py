class Foo:
    def __call__(self, *args, **kwargs):
        print('call....')

    def test(self):  #
        print('test....')


t = Foo()  # 实例化类
t.test()  # 正常调用实例方法
t()  # 直接调用实例化之后的对象

print('\n--------------------\n')
class Fuck:
    def __init__(self, func):
        self.func = func

    def __call__(self, *args, **kwargs):
        import time
        start_time = time.time()
        res = self.func(*args, **kwargs)
        end_time = time.time()
        print('函数在运行 "%s" 运行费时:  %s' % (self.func.__name__,
                                                    (end_time - start_time)))
        return res


@Fuck
def run(name):
    import time
    time.sleep(1)
    return 'sb_%s' % name


print(run('hyf'))

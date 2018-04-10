from threading import Lock

class LockedSet(set):
    """A set where add(), remove(), and 'in' operator are thread-safe"""

    def __init__(self, *args, **kwargs):
        self._lock = Lock()
        super(LockedSet, self).__init__(*args, **kwargs)

    def add(self, elem):
        with self._lock:
            super(LockedSet, self).add(elem)

    def remove(self, elem):
        with self._lock:
            super(LockedSet, self).remove(elem)

    def __contains__(self, elem):
        with self._lock:
            super(LockedSet, self).__contains__(elem)


# ---------------------------------------------------------

def lock(method):

  def newmethod(self, *args, **kwargs):
      with self._lock:
          return method(self, *args, **kwargs)
  return newmethod

class DecoratorLockedSet(set):
  def __init__(self, *args, **kwargs):
      self._lock = Lock()
      super(DecoratorLockedSet, self).__init__(*args, **kwargs)

  @locked_method
  def add(self, *args, **kwargs):
      return super(DecoratorLockedSet, self).add(elem)

  @locked_method
  def remove(self, *args, **kwargs):
      return super(DecoratorLockedSet, self).remove(elem)


# ----------------------------
"""
类的修饰符
类修饰将我们前面的示例更进一步，而不是保护单个方法，我们可以保护类中的每个函数，以便所有调用都以线程安全的方式进行。

在下面的示例中，我们将研究如何实现类decorator函数。开始时的lock_class函数包含方法和锁工厂的列表，并返回一个lambda函数，该函数接受在decorator和lockfactory中指定的方法名称。

这调用make_threadsafe，它初始化了我们的传入类的实例，然后它定义了一个新的构造函数，它也调用self。_lock = lockfactory()。这个make_threadsafe函数然后遍历方法名中的所有方法，并使用lock_method函数锁定每个方法。

这是一种干净且简单的方法，可以将线程安全添加到整个类中，同时还可以选择我们希望锁定的函数:
"""

from threading import Lock


def lock_class(methodnames, lockfactory):
    return lambda cls: make_threadsafe(cls,
                                       methodnames, lockfactory)


def lock_method(method):
    if getattr(method, '__is_locked', False):
        raise TypeError("Method %r is already locked!" % method)


def locked_method(self, *arg, **kwarg):
    with self._lock:
        return method(self, *arg, **kwarg)
    locked_method.__name__ = '%s(%s)' %


    ('lock_method', method.__name__)
    locked_method.__is_locked = True
    return locked_method


def make_threadsafe(cls, methodnames,
                    lockfactory):
    init = cls.__init__

    def newinit(self, *arg, **kwarg):
        init(self, *arg, **kwarg)
        self._lock = lockfactory()

    cls.__init__ = newinit

    for methodname in methodnames:
        oldmethod = getattr(cls, methodname)
        newmethod = lock_method(oldmethod)
        setattr(cls, methodname, newmethod)
    return cls


@lock_class(['add', 'remove'], Lock)
class ClassDecoratorLockedSet(set):

    @lock_method  # if you double-lock a method, a TypeError is raised
    def lockedMethod(self):
        print("This section of our code would be thread safe")
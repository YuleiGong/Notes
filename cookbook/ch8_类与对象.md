# 类与对象
## 改变对象字符串显示
* __repr__() 会返回一个实例代码的表示形式
* __str__() 使用str()或print()会调用此函数
* 如果没有使用__str__就会使用__repr__代替输出
* ({0.x!r}, {0.y!r}) 0指代了self本身

```python
class Pair:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return 'Pair({0.x!r}, {0.y!r})'.format(self)

    def __str__(self):
        return '({0.x!r}, {0.y!r})'.format(self)
out:
p = Pair(3, 4)
>>> p #repr
Pair(3, 4) 
>>> print (p) #str
(3, 4)
```

## 通过format()函数自定义字符串的格式化

```python
_formats = {
    'ymd' : '{d.year}-{d.month}-{d.day}',
    'mdy' : '{d.month}/{d.day}/{d.year}',
    'dmy' : '{d.day}/{d.month}/{d.year}'
}

class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    def __format__(self, code):
        if code == '':
            code = 'ymd'
        fmt = _formats[code]
        return fmt.format(d=self)
out:
d = date(2012, 12, 21)
>>> format(d)
'2012-12-21'
>>> format(d, 'mdy')
'12/21/2012'
```

## 让对象支持上下文管理器
* 当出现with语句,__entry__方法会被处罚,返回值会被赋值给 as 后的s,结束后执行__exit__触发清除工作
* 即使with中发生了一次,都会执行最后的清除工作,异常会被传递给__exit__()函数
* __exit__可以自己处理异常 exc_ty:异常类型 exc_val:异常值 tb:追溯信息

```python
from socket import socket, AF_INET, SOCK_STREAM
class LazyConnection:
    def __init__(self, address, family=AF_INET, type=SOCK_STREAM):
        self.address = address
        self.family = family
        self.type = type
        self.sock = None

    def __enter__(self):
        if self.sock is not None:
            raise RuntimeError('Already connected')
        self.sock = socket(self.family, self.type)
        self.sock.connect(self.address)
        return self.sock

    def __exit__(self, exc_ty, exc_val, tb):
        self.sock.close()
        self.sock = None

if __name__ == '__main__':
    from functools import partial
    conn = LazyConnection(('www.python.org', 80))
    with conn as s:
        s.send(b'GET /index.html HTTP/1.0\r\n')
        s.send(b'Host: www.python.org\r\n')
        s.send(b'\r\n')
        resp = b''.join(iter(partial(s.recv, 8192), b''))
        print (resp)

```

* 嵌套使用with 并允许多个连接,使用一个list来存储连接

```python
from socket import socket, AF_INET, SOCK_STREAM

class LazyConnection:
    def __init__(self, address, family=AF_INET, type=SOCK_STREAM):
        self.address = address
        self.family = family
        self.type = type
        self.connections = list()

    def __enter__(self):
        sock = socket(self.family, self.type)
        sock.connect(self.address)
        self.connections.append(sock)
        return sock

    def __exit__(self, exc_ty, exc_val, tb):
        self.connections.pop().close()

if __name__ == '__main__':
    from functools import partial
    conn = LazyConnection(('www.python.org', 80))
    with conn as s1:
        pass
        with conn as s2:
            pass
```
## 在类中对属性进行封装
* python 不会去依赖语言特效和封装数据,而是遵循一定的属性和命名规则来达到这个效果,使用***_***单下划线来达到这个效果(类内部实现)

```python
class A:
    def __init__(self):
        self._internal = 0
        self.public = 1

    def public_method(self):
        pass
    def _initernal_method(self):
        pass
```
* 使用***(__)***双下划线命名变量或方法,函数或变量会变得不可访问,实际上,变量或函数只是被重命名了
* 变量含函数会被重命名为_B__private _B__private_method 所以在外部无法访问
* 在继承中,使用这种方法变量或方法不会被覆盖,本例中子类的属性重命名为_C__private _C__private_method
* 大多数情况非公共名称以单下划线开头。如果使用继承,并且有些内部属性应该在子类中隐藏起来，使用双下划线方案。

```python
class B:
    def __init__(self):
        self.__private = 0

    def __private_method(self):
        pass
    def public_mechod(self):
        pass
        self.__private_method()

class C(B):
    def __init__(self):
        super().__init__()
        self.__private = 1

    def __private_method(self):
        pass
```
* 定义的变量和某个保留关键字冲突,使用单下划线后缀

```python
lambda_ = 2.0
```
## 创建可管理的属性
* 简单方法可以给方法定义一个property属性,想当于get方法。
* 只有first_name属性被定义,setter 和 deleter才会被创建

```python
class Person:
    def __init__(self, first_name):
        self.first_name = first_name

    @property
    def first_name(self):
        return self._first_name

    @first_name.setter
    def first_name(self, value):
        if not isinstance(value, str):
            raise TypeError('Exceped a string')
        self._first_name = value

    @first_name.deleter
    def first_name(name):
        raise AttributeError("Can't delete attribute")
```
* 访问时会自动触发setter getter deleter
* 实际变量存储在_first_name中,在调用__init__ 方法时会自动调用setter实施检查
* 不要在没有任何检查的代码里使用property

```python
>>> a = Person('Guido')
>>> a.first_name
'Guido'
>>> a.first_name = 'Bob'
>>> a.first_name
'Bob'
```

## 调用父类的方法
* 使用super方法调用

```python
class A:
    def spam(self):
        print ('A.spam')

class B(A):
    def spam(self):
        print ('B.spam')
        super().spam()
out:
>>> b = B()
>>> b.spam()
B.spam
A.spam
```
* 使用super初始化父类属性

```python
class A:
    def __init__(self):
        self.x = 0

class B(A):
    def __init__(self):
        super().__init__()
        self.y = 1
```

## 简化数据结构的初始化
* 很多class只用作存储数据,避免写很多的__init__()函数
* 在基类中使用一个公共的构造函数,使用setattr为子类的参数赋值

```python
import math
class Structure1:
    _fields = []
    def __init__(self, *args):
        if len(args) != len(self._fields):
            raise TypeError('Expected {} arguments'.format(len(self._fields)))
        for name, value in zip(self._fields, args):
            #self.name = value
            setattr(self, name, value)

class Stock(Structure1):
    _fields = ['name', 'shares', 'price']

class Point(Structure1):
    _fields = ['x', 'y']

class Circle(Structure1):
    _fields = ['radius']

    def area(self):
        return math.pi * self.radius ** 2
out:
>>> s = Stock('ACME', 50, 91.1)
>>> p = Point(2, 3)
>>> c = Circle('ACME', 50)
Traceback (most recent call last):
  File "<input>", line 1, in <module>
    c = Circle('ACME', 50)
  File "/Users/gongyulei/code_example/python_example/cookbook/ch8/ch8_11.py", line 11, i
n __init__
    raise TypeError('Expected {} arguments'.format(len(self._fields)))
TypeError: Expected 1 arguments
>>> s.shares
50
```
* 支持关键字参数作为传入参数

```python
class Structure2:
    _fields = []
    def __init__(self, *args, **kwargs):
        if len(args) > len(self._fields):
            raise TypeError('Expected {} arguments'.format(len(self._fields)))
        for name, value in zip(self._fields, args):
            #self.name = value
            setattr(self, name, value)

        for name in self._fields[len(args):]:
            setattr(self, name, kwargs.pop(name))

        if kwargs:
            raise TypeError('Invalid argument(s):{}'.format(',', join(kwargs)))

if __name__ == '__main__':
    class Stock(Structure2):
        _fields = ['name', 'shares', 'price']
    s1 = Stock('ACME', 50, 91.1)
    s2 = Stock('ACME', shares=91.1, price=91.1)
```
* 在构造子类的实例时加入在fields中不存在的属性

```python
class Structure3:
    _fields = []
    def __init__(self, *args, **kwargs):
        if len(args) != len(self._fields):
            raise TypeError('Expected {} arguments'.format(len(self._fields)))
        for name, value in zip(self._fields, args):
            #self.name = value
            setattr(self, name, value)

        extra_args = kwargs.keys() - self._fields
        for name in extra_args:
            setattr(self, name, kwargs.pop(name))

        if kwargs:
            raise TypeError('Invalid argument(s):{}'.format(',', join(kwargs)))

if __name__ == '__main__':
    class Stock(Structure3):
            _fields = ['name', 'shares', 'price']

        s1 = Stock('ACME', 50, 91.1, date='8/2/2012')

```
## 实现自定义容器
* 自定义类来模拟容器的功能,比如字典,列表。
* 要实现的类继承至collections,并实现部分方法
* 可以查看需要实现什么方法:len getitem

```python
>>> import collections
>>> collections.Sequence()
Traceback (most recent call last):
  File "<input>", line 1, in <module>
    collections.Sequence()
TypeError: Can't instantiate abstract class Sequence with abstract methods __getitem__, 
__len__
```
* 实现实现了len 和 getitem方法,完成对class的迭代

```python
import bisect
import collections

class SortedItems(collections.Sequence):
    def __init__(self, initila=None):
        self._items = sorted(initila) if initila is not None else []

    def __getitem__(self, index):
        return self._items[index]

    def __len__(self):
        return len(self._items)

    def add(self, item):
        #在插入的时候进行排序
        bisect.insort(self._items, item)

if __name__ == '__main__':
    items = SortedItems([5, 1, 3])
    print (list(items))
    print (items[0], items[-1])
    items.add(2)
    print (list(items))
out:
[1, 3, 5]             
1 5                   
[1, 2, 3, 5]  
```
* 可以使用isinstance实现类型检查

```python
>>> from ch8_14 import SortedItems
>>> import collections
>>> items = SortedItems()
>>> isinstance(items, collections.Sequence)
True
```
* collections 中很多抽象类会为一些常见容器操作提供默认的实现,只需要实现部分方法即可

```python
class Items(collections.MutableSequence):
    def __init__(self, initial=None):
        self._items = list(initial) if initial is not None else []

    def __getitem__(self, index):
        print ('Getting:', index)
        return self._items[index]

    def  __setitem__(self, index, value):
        print ('Setting:', index, value)
        self._itmes[index] = value

    def __delitem__(self, index):
        print ('Deleting', index)
        del self._itmes[index]

    def insert(self, index, value):
        print ('Inserting', index, value)
        self._items.insert(index, value)

    def __len__(self):
        print ('Len')
        return len(self._items)

out:
>>> a = Items([1, 2, 3])
>>> len(a)
Len
3
>>> a.append(4)
Len
Inserting 3 4
>>> a[0]
Getting: 0
1
```

## 属性的代理访问
* 某个实例的属性访问代理到另一个实例中去,在另一个类中使用
* 简单的代理

```python
class A:
    def spam(self, x):
        pass
    def foo(self):
        pass

class B1:
    def __init__(self):
        self._a = A()

    def spam(self, x):
        return self._a.spam(x)

    def foo(self):
        return self._a.foo()

    def bar(self):
        pass
```
* 大量方法需要代理,使用getattr,在调用抛出attribute异常属性不存在时,会自动调用A对象的相关方法

```python
class B2:
    def __init__(self):
        self._a = A()

    def bar(self):
        pass

    def __getattr__(self, name):
        #self._a.name
        return getattr(self._a, name)
out:
>>> from ch8_15 import B2
>>> b = B2()
>>> b.bar()
>>> b.spam(42)
```
* 实现代理模式
* __getattr__ 实际上是一个后备方法,proxy不存在会调用此函数
* __setattr__, __delattr__通过下划线区分代理实例和被代理实例的属性

```python
class Proxy:
    def __init__(self, obj):
        self._obj = obj

    def __getattr__(self, name):
        print ('getattr:', name)
        return getattr(self._obj, name)

    def __setattr__(self, name, value):
        if name.startswith('_'):
            super().__setattr__(name, value)
        else:
            print ('setattr:', name, value)
            setattr(self._obj, name, value)

    def __delattr__(self, name):
        if name.startswith('_'):
            super().__delattr__(name)
        else:
            print ('delattr:', name)
            setattr(self._obj, name)

class spam:
    def __init__(self, x):
        self.x = x

    def bar(self, y):
        print ('Spam.bar:', self.x, y)


if __name__ == '__main__':
    s = spam(2)
    p = Proxy(s)
    print (p.x)
    p.bar(3)
    p.x = 37
```

## 在类中定义多个构造器(实现init)
* 类方法的一个主要用途就是定义多个构造器。它接受一个 class 作为第一个参数(cls),返回一个实例

```python
import time
class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def today(cls):
        t = time.localtime()
        return cls(t.tm_year, t.tm_mon, t.tm_mday)
out:
>>> a = Date(2012, 12, 21)
>>> b = Date.today()
>>> b.day
30
```
## 不通过init方法初始化实例
* 需要批量创建实例,可以使用这种方法

```python
class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

if __name__ == '__main__':
    data = {'year':2012, 'month':8, 'day':29}
    d = Date.__new__(Date)
    for key, value in data.items():
        setattr(d, key, value)
```
## 实现状态机,避免过多的条件判断
* 定义每个状态类的基类

```python
class ConnectionState:
    @staticmethod
    def read(conn):
        raise NotImplementedError()

    @staticmethod
    def write(conn, data):
        raise NotImplementedError()

    @staticmethod
    def open(conn):
        raise NotImplementedError()

    @staticmethod
    def close(conn):
        raise NotImplementedError()

```
* 每个状态抽象定义出一个子类,并定义各自的静态方法

```python
class ClosedConnectionState(ConnectionState):
    @staticmethod
    def read(conn):
        raise RuntimeError('Not open')

    @staticmethod
    def write(conn, data):
        raise RuntimeError('Not open')

    @staticmethod
    def open(conn):
        conn.new_state(OpenConnectionState)

    @staticmethod
    def close(conn):
        raise RuntimeError('Already closed')

class OpenConnectionState(ConnectionState):
    @staticmethod
    def read(conn):
        print ('reading')

    @staticmethod
    def write(conn, data):
        print ('writing')

    @staticmethod
    def open(conn):
        raise RuntimeError('Already open')

    @staticmethod
    def close(conn):
        conn.new_state(ClosedConnectionState)

```
* 所有的状态实例存储在Connection中,通过动态的改变该实例状态,调用相关方法,自动实现状态控制

```python
class Connection:

    def __init__(self):
        self.new_state(ClosedConnectionState)

    def new_state(self, newstate):
        self._state = newstate

    def read(self):
        return self._state.read(self)

    def write(self, data):
        return self._state.write(self, data)

    def open(self):
        return self._state.open(self)

    def close(self):
        return self._state.close(self)


```
* test

```sh
>>> c = Connection()
>>> c._state
<class 'ch8_19.ClosedConnectionState'>
>>> c.read()
Traceback (most recent call last):
  File "<input>", line 1, in <module>
    c.read()
  File "/Users/gongyulei/code_example/python_example/cookbook/ch8/ch8_19.py", line 16, i
n read
    return self._state.read(self)
  File "/Users/gongyulei/code_example/python_example/cookbook/ch8/ch8_19.py", line 47, i
n read
    raise RuntimeError('Not open')
RuntimeError: Not open
>>> c.open()
>>> c._state
<class 'ch8_19.OpenConnectionState'>
>>> c.write('hello')
writing
>>> c.close()
>>> c._state
<class 'ch8_19.ClosedConnectionState'>
```
##  通过字符串调用对象的方法
* 使用getattr

```python
import math
class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def distance(self, x, y):
        return math.hypot(self.x - x, self.y - y)
out:
>>> p = Point(2, 3)
>>> d = getattr(p, 'distance')(0, 0)
>>> d
3.605551275463989
>>> d = getattr(p, 'distance')
>>> d
<bound method Point.distance of <ch8_20.Point object at 0x109815438>>
```

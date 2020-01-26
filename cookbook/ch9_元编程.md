# 元编程
## 在函数上添加装饰器
* 在函数上添加包装器,增加额外的操作
* 一个装饰器就是一个函数,接收一个函数作为参数并返回一个新的函数,在函数内部使用args 和 kwargs 来接收任意参数
* 在定义装饰器函数的时候使用@wraps来保留函数的元信息(函数名字,注释等)
* 使用@语法等效与

```python
def countdown(n):
    pass
countdown = timethis(countdown)
```

```python
import time
from functools import wraps
def timethis(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        end = time.time()
        print (func.__name__, end-start)
        return result
    return wrapper

@timethis
def countdown(n):
    while n > 0:
        n -= 1

if __name__ == '__main__':
    countdown(100000)

```

## 解除装饰器
* 撤销已经作用的装饰器,直接访问未包装的函数
* 如果使用@wraps装饰的函数,直接使用__wrapped__访问原始函数
* 如果有多个装饰器作用于同一个函数,不能使用这种方法

```python
>>> orig_add = countdown.__wrapped__
>>> orig_add(1000)
```

## 定义一个带参数的装饰器
* 在decorate中处理传入的参数

```python
from functools import wraps
import logging

def logged(level, name=None, message=Nome):
    def decorate(func):
        logname = name if name else func.__module__
        log = logging.getLogger(logname)
        logmsg = message if message else func.__name__

        @wraps(func)
        def wrapper(*args, **kwargs):
            log.log(level, logmsg)
            return func(*args, *kwargs)
        return wrapper
    return decorate

@logged(logging.CRITICAL)
def add(x, y):
    return x + y

@logged(logging.CRITICAL)
def spam():
    print ('spam!')

if __name__ == '__main__':
    add(1, 2)
    fun = spam()
```

## 带可选参数的装饰器
* 可以自由的控制装饰器的形参
* @logged()表示什么参数都没传func=None 
* @logged表示传递了一个func = add add=logged(add),被装饰函数会被当做第一个参数传递给logged装饰器
* @logged(level=logging.CRITICAL, name='example')等价于spam = logged(level=logging.CRITICAL, name='example')(spam),第一次调用返回一个partial函数给spam(),此时再次调用改装饰器调用形式为partial()(spam),将func 传递给logged

```python
def logged(func=None,*,level=logging.DEBUG, name=None, message=None):
    if func is None:
        return partial(logged, level=level, name=name, message=message)

    logname = name if name else func.__module__
    log = logging.getLogger(__name__)
    logmsg = message if message else func.__name__

    @wraps(func)
    def wrapper(*args, **kwargs):
        log.log(level, logmsg)
        return func(*args, **kwargs)

    return wrapper

@logged()
def add(x, y):
    return x + y

@logged(level=logging.CRITICAL, name='example')
def spam():
    print ('Spam!')

if __name__ == '__main__':
    pass

```

## 利用装饰器强制函数上的类型检查
* 对函数参数进行强制类型检查


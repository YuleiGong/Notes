# 函数

## 可接受任意数量参数的函数
* 可以使用 __*__  使函数接收任意数量的位置参数, __*__ 只能是函数的最后一个位置参数
* 可以使用 __**__ 使函数接收任意数量的关键字参数 __**__ 只能是函数的最后一个参数

```python
def func(x, *args, **kwargs):
    print (x)
    print (args)
    print (kwargs)

if __name__ == '__main__':
    func(1, 2, 3, four=4)
out:
1                     
(2, 3)                
{'four': 4}  
```

## 只接受关键字参数的函数
* 提升可读性

```python
def recv(num, block=None):
    pass

msg = recv(1024, block=False)
```
## 定义有默认参数的函数
* 使用_no_value赋值给object()测试某个参数是否传递了进来
* 传递None 和 不传递参数是有差别的

```python
_no_value = object()

def spam(a, b=_no_value):
    if b is _no_value:
        print ('No b value supplied')
    print (a)

if __name__ == '__main__':
    spam(1)
    spam(1, 2)
    spam(1, None)
out:
No b value supplied   
1                     
1                     
1    
```
* 默认参数的值应该是不可变的对象(None, True, Flase,数字或字符串)
* 使用可变对象传值,可能会无意间修改传递的默认值

```python
ERROR!
def spam(a, b=[]):
    print (b)
    return b
out:
>>> x = spam(1)
[]
>>> x.append(99)
>>> x.append('Yow!')
>>> x
[99, 'Yow!']
>>> spam(1)
[99, 'Yow!']
```
* 在检查None的时候使用is 操作符否则0, '',都会被检测为__None__

```python
def spam(a, b=None):
    #Error
    if not b:
        pass
    #correct
    if b is None:
        pass
```
## 定义匿名或内联函数(lambda)
* lambda 只能指定单个表达式,他的值就是最后的返回值

```pyhton
>>> add = lambda x, y:x+y
>>> add(2, 3)
5
>>> add('hell', 'o')
'hello'

#根据名的来进行排序
>>> names = ['David Beazley', 'Brian Jones','Raymond Hettinger', 'Ned Batchelder']
>>> sorted(names, key=lambda name:name.split()[-1].lower())
['Ned Batchelder', 'David Beazley', 'Raymond Hettinger', 'Brian Jones']
```

## 匿名函数捕获变量值
* lambda定义一个匿名函数,定义时捕获到某些变量值
* lambda中的x是一个自由变量,在运行时绑定

```python
>>> x = 10
>>> a = lambda y:x+y
>>> x = 20
>>> a(10)
30
```
* 在lambda中绑定默认值 __main()__,funcs函数列表的n值都不同

```python
def main():
    #根据lambda 生成函数
    funcs = [lambda x, n=n:x+n for n in range(5)]
    for f in funcs:
        print (f(0))

def main1():
    funcs = [lambda x :x+n for n in range(5)]
    for f in funcs:
        print (f(0))

out:
0,1,2,3,4
```
## 减少函数的参数个数
* 使用functools.partial()减少函数的调用参数值
* partial会返回一个新的函数对象,新的函数会合并已经赋值的参数,最后传递给原始函数

```python
def spam(a, b, c, d):
    print (a, b, c, d)
>>> s1 = partial(spam, d=42)
>>> s1(1, 2, 3)
1 2 3 42
```

## 将单方法的类转换为函数
* class 中除了__init__()外,只定义了一个方法
* 使用类的原因是需要存储一些额外的值给template调用,比如本例中的template
* 在闭包中,一个闭包就是一个函数,闭包会记住自己在定义时的环境,在本例中opener()函数记住了template变量
* 给函数使用闭包比使用类更加简洁

```python
from urllib.request import urlopen
class UrlTemplate:
    def __init__(self, template):
        self.template = template

    def open(self, **kwargs):
        return urlopen(self.template.format_map(kwargs))

def urltemplate(template):
    def opener(**kwargs):
        return urlopen(template.format_map(kwargs))
    return opener

if __name__ == '__main__':
    yahoo = UrlTemplate('http://finance.yahoo.com/d/quotes.csv?s={names}&f={fields}')
    for line in yahoo.open(names='IBM,AAPL,FB', fields='sl1c1v'):
        print (line.decode('utf-8'))


    yahoo = urltemplate('http://finance.yahoo.com/d/quotes.csv?s={names}&f={fields}')
    for line in yahoo(names='IBM,AAPL,FB', fields='sl1c1v'):
        print (line.decode('utf-8'))
```
## 带额外状态信息的回调函数
* 为了让回调函数访问外部信息，可以给回调函数绑定一个额外的方法,通过访问类的所属变量来达到目的

```python
def apply_async(func, args, *, callback):
    result = func(*args)
    callback(result)

def print_result(result):
    print ('Got:', result)

def add(x, y):
    return x + y

class ResultHandler:
    def __init__(self):
        self.sequence = 0

    def handler(self, result):
        self.sequence += 1
        print ('[{}] Got:{}'.format(self.sequence, result))

if __name__ == '__main__':
    apply_async(add, (2, 3), callback=print_result)
    #通过类方法来访问外部信息
    r = ResultHandler()
    apply_async(add, (2, 3), callback=r.handler)
    apply_async(add, ('hello', 'world'), callback=r.handler)
```
* 可以使用闭包来代替类方法来捕获状态值
* __nonlocal__ 声明可以改变闭包中的参数值,这里是sequence += 1

```python
def apply_async(func, args, *, callback):
    result = func(*args)
    callback(result)

def add(x, y):
    return x + y

def make_handler():
    sequence = 0
    def handler(result):
        nonlocal sequence
        sequence += 1
        print('[{}] Got: {}'.format(sequence, result))
    return handler

if __name__ == '__main__':
    #使用闭包捕获状态值
    handler = make_handler()
    apply_async(add, (2, 3), callback=handler)
    apply_async(add, ('hello', 'world'), callback=handler)
```

## 访问闭包中定义的变量,并且可以修改
* 闭包的内部变量对于外界是完全隐藏的,但可以编写访问函数并将函数绑定到闭包上实现
* 函数属性允许我们用一种很简单的方式将访问方法绑定到闭包函数上，这个跟实例方法很像

```python
def sample():
    n = 0
    def func():
        print ('n=', n)

    def get_n():
        return n

    def set_n(value):
        nonlocal n
        n = value

    func.get_n = get_n
    func.set_n = set_n
    return func
out:
>>> f = sample()
>>> f()
n= 0
>>> f.set_n(10)
>>> f()
n= 10
>>> f.get_n()
10
```

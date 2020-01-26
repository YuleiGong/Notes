# 迭代器与生成器

## 手动遍历迭代器
* 不使用for循环手动遍历迭代器

```python
def manual_iter():
    '''
    手动遍历迭代器
    '''
    with open('/etc/passwd') as f:
        try:
            while True:
                line = next(f)
                print (line)
        #通过StopIteration异常来标记结尾
        except StopIteration:
            pass

def manual_iter_1():
    with open('/etc/passwd') as f:
        while True:
            line = next(f, None)
            if line is None:
                break
            print (line)
```
## 代理迭代
* 在一个新的容器(class)中迭代对象

```python
class Node:
    def __init__(self, value):
        self._value = value
        #存储Node对象
        self._children = []

    def __repr__(self):
        return 'Node({!r})'.format(self._value)

    def add_child(self, node):
        self._children.append(node)

    def __iter__(self):
        '''
        传递迭代请求给内部的_children,实现了一个__next()__方法的迭代器对象
        '''
        return iter(self._children)

if __name__ == '__main__':
    root = Node(0)
    child1 = Node(1)
    child2 = Node(2)
    root.add_child(child1)
    root.add_child(child2)
    for ch in root:
        print (ch)

```
## 使用生成器创建新的迭代模式
* 生成器只能用于迭代操做，使用for循环迭代生成器会自动处理StopIteration异常

```python
def frange(start, stop, increment):
    x = start
    while x < stop:
        yield x
        x += increment

def countdown(n):
    print ('Starting to count from', n)
    while n > 0:
        yield n
        n -= 1
    print ('Done!')
```
## 实现迭代器协议
* 构建一个支持迭代操作的自定义对象

```python
class Node:
    def __init__(self, value):
        self._value = value
        #存储Node对象
        self._children = []

    def __repr__(self):
        return 'Node({!r})'.format(self._value)

    def add_child(self, node):
        self._children.append(node)

    def __iter__(self):
        #传递迭代请求给内部的_children
        return iter(self._children)
        #return self._children.__iter__()

    def depth_first(self):
        #返回自己本身的一个迭代
        yield self
        for c in self:
            #调用迭代本身的函数
            yield from c.depth_first()

if __name__ == '__main__':
    #根
    root = Node(0)
    child1 = Node(1)
    child2 = Node(2)
    root.add_child(child1)
    root.add_child(child2)
    child1.add_child(Node(3))
    child1.add_child(Node(4))
    child2.add_child(Node(5))
    for ch in root.depth_first():
        print (ch)

```
## 反向迭代
* 使用内置的reversed()
* 反向迭代仅仅当对象的大小可预先确定或者对象实现了 __reversed__() 的特殊方法时才能生效

```python
>>> a = [1, 2, 3, 4]
>>> for x in reversed(a):
...     print (x)
4
3
2
1
```
* 通过__reversed__自定义类的反向迭代

```python
class Countdown:

    def __init__(self, start):
        self.start = start

    def __iter__(self):
        n = self.start
        while n > 0:
            yield n
            n -= 1

    def __reversed__(self):
        n = 1
        while n <= self.start:
            yield n
            n += 1

if __name__ == '__main__':
    #调用自己实现的__reversed__方法
    for rr in reversed(Countdown(30)):
        print (rr)
    for rr in Countdown(30):
        print (rr)
```
## 带有外部状态的生成器函数
* 定义一个生成器函数，调用某个你想暴露给用户使用的外部状态值
* 通过类的__iter__函数实现

```python
class LineHistory:
    def __init__(self, lines, histlen=3):
        self.lines = lines
        self.history = deque(maxlen=histlen)

    def __iter__(self):
        #遍历并标记计数起始点
        for lineno, line in enumerate(self.lines, 1):
            self.history.append((lineno, line))
            yield line

    def clear(self):
        self.history.clear()

if __name__ == '__main__':
    with open('ch4_1.py') as f:
        lines = LineHistory(f)
        #__iter_函数返回的生成器 
        for line in lines:
            if 'python' in line:
                for lineno, hline in lines.history:
                    print ('{}:{}'.format(lineno, hline))
```
## 迭代器切片
* 得到一个由迭代器生成的切片对象
* 迭代器和生成器不能使用标准的切片操作，因为它们的长度事先我们并不知道(并且也没有实现索引)。 函数 islice() 返回一个可以生成指定元素的迭代器，它通过遍历并丢弃直到切片开始索引位置的所有元素。 然后才开始一个个的返回元素，并直到切片结束索引位置。
*  islice() 会消耗掉传入的迭代器中的数据

```python
import itertools
def count(n):
    while True:
        yield n
        n += 1

if __name__ == '__main__':
    c = count(0)
    for x in itertools.islice(c, 10, 20):
        print (x)
```
## 跳过可迭代对象的开始部分
* 跳过具体部分(开头是#的文件)

```python
from itertools import dropwhile
def open_passwd():
    with open('/etc/passwd') as f:
        for line in dropwhile(lambda line:line.startswith('#'), f):
            print (line)
```
* 如果明确知道跳过的个数

```python
>>> from itertools import islice
>>> items = ['a', 'b', 'c', 1, 2, 3]
>>> for x in islice(items, 3, None):
...     print (x)
1
2
3
```
## 排列组合迭代
* 迭代遍历一个集合中元素的所有可能的排列或组合
* permutations接受一个集合并产生一个元组序列，每个元组由集合中所有元素的一个可能排列组成

```python
>>> items = ['a', 'b', 'c']
>>> from itertools import permutations
>>> for p in permutations(items):
...     print (p)
('a', 'b', 'c')
('a', 'c', 'b')
('b', 'a', 'c')
('b', 'c', 'a')
('c', 'a', 'b')
('c', 'b', 'a')

```
* 指定长度的排列

```python
>>> for p in permutations(items, 2):
...     print (p)
('a', 'b')
('a', 'c')
('b', 'a')
('b', 'c')
('c', 'a')
('c', 'b')
```
* 指定长度的组合(不考虑排序)

```python
>>> from itertools import combinations
>>> for c in combinations(items, 3):
...     print (c)
('a', 'b', 'c')
>>> for c in combinations(items, 2):
...     print (c)
('a', 'b')
('a', 'c')
('b', 'c')
```
## 序列上索引值迭代
* enumerate()可以很好解决
* 指定一个start，改变索引开始的值

```python
>>> for id, val in enumerate(my_list, 1):
...     print (id, val)
...     
... 
1 a
2 b
3 c
```
* 在遍历文件时定位错误行数

```python
def pares_data(filename):
    with open(filename, 'r') as f:
        for lineno, line in enumerate(f, 1):
            fields = line.split()
            try:
                count = int(fields[0])
            except ValueError as e:
                print ('Line {}: Parse error:{}'.format(lineno, e))
```
* 跟踪某些单词在出现的函数，构造出一个字典

```python
def line_dict():
    from collections import defaultdict
    #初始化字典,value=list
    word_summary = defaultdict(list)
    with open('ch4_1.py', 'r') as f:
        lines = f.readlines()

    for idx, line in enumerate(lines):
        words = [w.strip().lower() for w in line.split()]
        for word in words:
            word_summary[word].append(idx)
    return word_summary
```
## 同时迭代多个序列
* 使用zip(), 迭代的次数和最短的序列一致

```python
>>> xpts = [1, 5, 4, 2, 10, 7]
>>> ypts = [101, 78, 37, 15, 52, 99]
>>> for x, y in zip(xpts, ypts):
...     print (x, y)
1 101
5 78
4 37
2 15
10 52
7 99
```
* 可以使用itertools.zip_longest()来填充没有zip迭代到的默认值

```python
>>> for i in zip_longest(a, b):
...     print (i)
(1, 'w')
(2, 'x')
(3, 'y')
(None, 'z')
>>> for i in zip_longest(a, b, fillvalue=0):
...     print (i)
(1, 'w')
(2, 'x')
(3, 'y')
(0, 'z')
```
* 如果zip的列表一个是key，一个是value,可以打包成字典

```
>>> headers = ['name', 'shares', 'price']
>>> values = ['ACME', 100, 490.1]
>>> s = dict(zip(headers, values))
>>> print (s) {'name': 'ACME', 'shares': 100, 'price': 490.1}
```
## 不同集合上元素的迭代
* 在不同的可迭代对象上循环变量(避免写重复的for循环)

```python
>>> from itertools import chain
>>> a = [1, 2, 3, 4]
>>> b = ['x', 'y', 'z']
>>> for x in chain(a, b):
...     print (x)
1
2
3
4
x
y
z
```
## 数据处理管道

```python
import os
import fnmatch
import gzip
import bz2
import re

def gen_find(filepat, top):
    #遍历当期文件夹的所有内容,返回当期文件夹路径,文件夹列表，文件列表
    for path, dirlist, filelist in os.walk(top):
        #找到匹配的日志,重新组装路径
        for name in fnamatch.filter(filelist, filepat):
            yield os.path.join(path, name)

def gen_opener(filenames):
    #每次返回一个打开的文件描述符
    for filename in filenames:
        if filename.endswith('.gz'):
            f = gzip.open(filename, 'tr')
        elif filename.endswith('.bz2'):
            f = bz2.open(filename, 'rt')
        else:
            f = open(filename, 'rt')
        yield f
        f.close()

def gen_concatenate(iterators):
    #拼接打开的内容
    for it in iterators:
        yield from it

def gen_grep(pattern, lines):
    pat = re.complate(pattern)
    for line in lines:
        if pat.search(line):
            yield line

if __name__ == '__main__':
    lognames = gen_find('access-log*', 'www')
    files = gen_opener(lognames)
    pylines = gen_grep('(?i)python', lines)
    for line in pylines:
        print (lines)
```
## 展开嵌套的序列
* 将一个多层嵌套序列展开为单层列表
* Iterable检查元素是否是可以迭代的,ignore_types防止迭代字符数据

```python
from collections import Iterable
def flatten(items, ignore_types=(str, bytes)):
    for x in items:
        if isinstance(x, Iterable) and not isinstance(x, ignore_types):
            yield from flatten(x)
        else:
            yield x

if __name__ == '__main__':
    items = [1, 2, [3, 4, [5, 6], 7], 8]
    for x in flatten(items):
        print (x)
```
## 顺序迭代合并后的序列
* 要求输入序列是经过排序的

```python
>>> import heapq
>>> a  = [1, 4, 7, 10]
>>> b = [2, 5, 6, 11]
>>> for c in heapq.merge(a, b):
...     print (c)
1
2
4
5
6
7
10
11
```






# 数据结构和算法
## 通过*表达式或者是一个用不到的变量来拆分,处理序列
* 只要对象恰好是可迭代的,就可以执行分解操作
* 使用 _ 符号用于__舍弃__序列中不需要的值

```python
>>> data = ['a', 'v', 'x']
>>> a, _, x = data
>>> a
'a'
>>> _
'v'
>>> x
'x'
>>> 
```
* 使用 * 表达式分解元素。

```python
def drop_first_last(grads):
 ''' 
    *号去除变量
    *号解压出的变量永远是列表
    '''
    first, *middle, last = grads
    return middle

```

```python
def split_star():
    '''
    *号切分串
    '''
    line = 'nobody:*:-2:-2:Unprivileged User:/var/empty:/usr/bin/false'
    uname, *fiedles, homedir, sh = line.split(':')
    print (uname)
    print (homedir)
    print (sh)
    print (fiedles) 
```

* *号实现递归
```python
def sum(items):
    '''
    利用*分解操作实现某个数组的递归加
    '''
    head, *tail = items
    return head + sum(tail) if tail else head
```
## 使用collenctions的dque可以构造一个队列
* 设置队列的长度,append满后可以自动移除最先入队的元素
```python
>>> q = deque(maxlen=5)
>>> q.append(1)
>>> q.append(2)
>>> q.append(3)
>>> q.append(4)
>>> q
deque([1, 2, 3, 4], maxlen=5)
>>> q.append(5)
>>> q
deque([1, 2, 3, 4, 5], maxlen=5)
>>> q.append(6)
>>> q
deque([2, 3, 4, 5, 6], maxlen=5)
```
* 可以指定append 和 pop的方向
```python
q.popleft()
q.appendleft()
```

## 使用heapq.nlargest 和 heapq.nsmallest 灵活找出集合中最大或者最小的N个元素
* 还可以使用lambda表达式执行复杂的查找
```python 
import heapq                                                     def test_heapq():
    nums = [1, 8, 2, 23, 7, -4, 18, 23, 42, 37, 2]
    portfolio = [
        {'name':'a', 'num':1},
        {'name':'b', 'num':2},
        {'name':'c', 'num':3},
        {'name':'d', 'num':4},
        {'name':'e', 'num':5}]
    print (heapq.nlargest(3, nums))
    print (heapq.nsmallest(3, nums))
    #使用key和lambda复杂的排序
    print (heapq.nsmallest(3, portfolio, key=lambda s:s['num']))
    print (heapq.nlargest(3, portfolio, key=lambda s:s['num']))
```
* 如果只是找出最大或最小,N=1,那么使用max() min()是最适合的

### 如果N相对于总数来说,N很小,使用堆是最适合的,每次Pop都能pop出最小的元素
```python
>>> import heapq
>>> num = [1, 2, 3, 4, 8, 5, 6]
>>> heapq.heapify(num)
>>> heapq.heappop(num)
1
>>> heapq.heappop(num)
2
>>> heapq.heappop(num)
3
```

## 使用defaultdict 可以自动初始化一对多字典的初始值
```python
>>> from collections import defaultdict
>>> d = defaultdict(list)
>>> d['a'].append(1)
>>> d['a'].append(2)
>>> d['b'].append(3)
>>> d
defaultdict(<class 'list'>, {'b': [3], 'a': [1, 2]})
>>> d[a]
[1, 2]
>>> d = defaultdict(set)
>>> d['a'].add(1)
>>> d['a'].add(1)
>>> d['a'].add(2)
>>> d['b'].add(3)
>>> d
defaultdict(<class 'set'>, {'b': {3}, 'a': {1, 2}})
>>> d['a']
{1, 2}
>>> 
```

## 控制字典迭代的顺序
* 使用 OrderedDict控制字典迭代顺序
```python
>>> from collections import OrderedDict
>>> d = OrderedDict()
>>> d['a'] = 1
>>> d['b'] = 2
>>> d['c'] = 3
>>> for key in d:
...     print (key, d[key])
a 1
b 2
c 3
```
* 在构建json时也能控制顺序
* OrderedDict内部维护了一个双向列表,数据量大的时候需要考虑内存开销
```python
>>> import json
>>> json.dumps(d)
'{"a": 1, "b": 2, "c": 3}'
```

## 字典的排序,极值
* 使用zip对反转键值对,使用max或者min找出最大最小值。zip出来的迭代器只能被消费一次,不能重复使用
* min max 对元组进行比较,会先对values进行比较,而后才是key
```python
>>> price = {'ACME':45.23, 'AAPL':712.78, 'FB':11}
>>> price_zip = zip(price.values(), price.keys())
>>> min(price_zip)
(11, 'FB')
>>> max(price_zip)
Traceback (most recent call last):
  File "<input>", line 1, in <module>
    max(price_zip)
ValueError: max() arg is an empty sequence
```
* 不使用zip,在min max 时候指定key.内部排序在遍历dict时使用pirce[k]即value进行排序返回key
```python
>>> price = {'ACME':45.23, 'AAPL':712.78, 'FB':11, 'IBM':205.55}
>>> min(price, key=lambda k:price[k])
'FB'
```
## 在两个字典中寻找相同点
```python
>>> a = {'x':1, 'y':2, 'z':3}
>>> b = {'w':10, 'x':11, 'y':2}
# Find keys in common
>>> a.keys() & b.keys()
{'y', 'x'}
# Find keys in a that are not in b
>>> a.keys() - b.keys()
{'z'}
# Find key,value pairs in a that not in b
>>> a.items() - b.items()
{('x', 1), ('z', 3)}
```
## 删除序列相同的元素并保持元素顺序
* yield 返回了一个生成器
* 不考虑保持元素顺序,可以使用set去除相同的元素
```python
def dedupe(items):
    seen = set()
    for item in items:
        if item not in seen:
            yield item
            seen.add(item)
if __name__ == '__main__':
    a = [1, 5, 2, 1, 9, 1, 5, 10] 
    print (dedupe(a))
    for i in dedupe(a):
        print (i) 
```
## 命名切片slice
* 避免使用硬编码切片
```python
>>> a = slice(0, 10, 2)
>>> c = 'dsfdffdfdfdfd'
>>> c[a]
'dffdd'
>>> a = slice(0, 10)
>>> c[a]
'dsfdffdfdf'
```
* 调用indices可以返回一个切片序列的元组,用于迭代切片
* 使用*将元组放入range
```python
>>> s = 'HelloWorld'
>>> for i in range(*slice(5, 10, 2).indices(10)):
...     print (i)
...     print (s[i])
5
W
7
r
9
d
```
## 统计速序中出现次数最多的元素
```python
>>> words = [
...     'look', 'into', 'my', 'eyes', 'look', 'into', 'my', 'eyes',
...     'the', 'eyes', 'the', 'eyes', 'the', 'eyes', 'not', 'around', 'the',
...     'eyes', "don't", 'look', 'around', 'the', 'eyes', 'look', 'into',
...     'my', 'eyes', "you're", 'under'
... ]
>>> from collections import Counter
>>> word_counts = Counter(words)
>>> top_three = word_counts.most_common(3)
>>> print (top_three)
[('eyes', 8), ('the', 5), ('look', 4)]
>>> top_three = word_counts.most_common(1)
>>> print (top_three)
[('eyes', 8)]
```
* 一个Counter对象就是一个字典
```python
>>> word_counts['not']
1
>>> word_counts['eyes']
8
```
* Counter也可以进行数学运算,在制表和计数数据的场合很有用
```python
>>> a = Counter(words)
>>> a
Counter({'eyes': 8, 'the': 5, 'look': 4, 'into': 3, 'my': 3, 'around': 2, 'not
': 1, "don't": 1, "you're": 1, 'under': 1})
>>> morewords = ['eyes', 'eyes', 'look']
>>> b = Counter(morewords)
>>> c = a + b
>>> c
Counter({'eyes': 10, 'look': 5, 'the': 5, 'into': 3, 'my': 3, 'around': 2, 'no
t': 1, "don't": 1, "you're": 1, 'under': 1})
>>> d = a - b
>>> d
Counter({'eyes': 6, 'the': 5, 'look': 3, 'into': 3, 'my': 3, 'around': 2, 'not
': 1, "don't": 1, "you're": 1, 'under': 1})
```
## 通过某个关键字排序字典序列
```python
>>> rows = [
...     {'fname': 'Brian', 'lname': 'Jones', 'uid': 1003},
...     {'fname': 'David', 'lname': 'Beazley', 'uid': 1002},
...     {'fname': 'John', 'lname': 'Cleese', 'uid': 1001},
...     {'fname': 'Big', 'lname': 'Jones', 'uid': 1004}
... ]
>>> from operator import itemgetter
>>> itemgetter('fname')
operator.itemgetter('fname')
>>> print (sorted(rows, key=itemgetter('fname')))
[{'fname': 'Big', 'lname': 'Jones', 'uid': 1004}, {'fname': 'Brian', 'lname':'Jones', 'uid': 1003}, {'fname': 'David', 'lname': 'Beazley', 'uid': 1002}, {'fname': 'John', 'lname': 'Cleese', 'uid': 1001}]
>>> print (sorted(rows, key=itemgetter('lname', 'fname')))
[{'fname': 'David', 'lname': 'Beazley', 'uid': 1002}, {'fname': 'John', 'lname': 'Cleese', 'uid': 1001}, {'fname': 'Big', 'lname': 'Jones', 'uid': 1004}, {'fname': 'Brian', 'lname': 'Jones', 'uid': 1003}]
```
## 排序不支持原生比较的class
* lambda 和 attregetter 都可以使用
* 如果要对比多个字段,需要使用attrgetter
```python
from operator import attrgetter
class User(object):

    def __init__(self, user_id):
        self.user_id = user_id

    def __repr__(self):
        return 'User({})'.format(self.user_id)
def sort_notcompare():
    users = [User(23), User(3), User(99)]
    print(users)
    print (sorted(users, key=lambda u:u.user_id))
    print (sorted(users, key=attrgetter('user_id')))
```
## 根据dict字段将记录分组，可以分组迭代
* 先根据date字段对rows排序
* 排序后groupby分组，会返回分组的字段值和分组内容的迭代器
```python
from operator import itemgetter
from itertools import groupby
def group_by(rows):
    #根据date对dict排序
    rows.sort(key=itemgetter('date'))
    for date, items in groupby(rows, key=itemgetter('date')):
        print (date)
        #生成迭代器
        for i in items:
            print (i) 

if __name__ == '__main__':
    rows = [ 
        {'address': '5412 N CLARK', 'date': '07/01/2012'},
        {'address': '5148 N CLARK', 'date': '07/04/2012'},
        {'address': '5800 E 58TH', 'date': '07/02/2012'},
        {'address': '2122 N CLARK', 'date': '07/03/2012'},
        {'address': '5645 N RAVENSWOOD', 'date': '07/02/2012'},
        {'address': '1060 W ADDISON', 'date': '07/02/2012'},
        {'address': '4801 N BROADWAY', 'date': '07/01/2012'},
        {'address': '1039 W GRANVILLE', 'date': '07/04/2012'},
    ]   
    group_by(rows)
```
## 筛选序列元素
* 通过列表推导式来筛选
```python
>>> mylist = [1, 4, -5, 10, -7, 2, 3, -1]
>>> [n for n in mylist if n > 0 ]
[1, 4, 10, 2, 3]
>>> 
```
* 不仅仅是筛选,还要替换不合规则的值
```python
>>> mylist = [1, 4, -5, 10, -7, 2, 3, -1]
>>> chip_neg = [n if n > 0 else 0 for n in mylist]
>>> chip_neg
[1, 4, 0, 10, 0, 2, 3, 0]
```
* 如果筛选数据较多通过生成器来筛选
```python
>>> mylist = [1, 4, -5, 10, -7, 2, 3, -1]
>>> pos = (n for n in mylist if n > 0 )
>>> pos
<generator object <genexpr> at 0x10e5eab48>
>>> for i in pos:
...     print(i)
1
4
10
2
3
```
* 如果筛选过程复杂，涉及异常处理，可以将筛选过程放在函数中，通过 __filter__ 处理，此函数会将列表里面的内容依次作用于函数，根据True和False来决定是否保留
* __filter__ 返回一个迭代器
```python
def is_int(val):
    try:
        x = int(val)
        return True
    except ValueError:
        return False
if __name__ == '__main__':
    values = ['1', '2', '-3', '-', 'N/A', '5']
    ivals = list(filter(is_int, values))
    print (ivals)
```
* __itertools.compress()__ 筛选
* __compress__ 返回一个迭代器，需要传入筛选列表的布尔表达式，此函数会筛选出True的值，常用于把一个序列的值施加到另一个序列上
```python
>>> addresses = [
...     '5412 N CLARK',
...     '5148 N CLARK',
...     '5800 E 58TH',
...     '2122 N CLARK',
...     '5645 N RAVENSWOOD',
...     '1060 W ADDISON',
...     '4801 N BROADWAY',
...     '1039 W GRANVILLE',
... ]
>>> counts = [ 0, 3, 10, 4, 1, 7, 6, 1]
>>> from itertools import compress
>>> more5 = [n>5 for n in counts]
>>> more5
[False, False, True, False, False, True, True, False]
>>> list(compress(addresses, more5))
['5800 E 58TH', '1060 W ADDISON', '4801 N BROADWAY']
```
## 字典中提取子集
* 使用字典推导式
```python
>>> prices = {
...     'ACME': 45.23,
...     'AAPL': 612.78,
...     'IBM': 205.55,
...     'HPQ': 37.20,
...     'FB': 10.75
... }
>>> p1 = {key:value for k, v in prices.items() if v > 200 }
>>> p1 = {k:v for k, v in prices.items() if v > 200 }
>>> p1
{'AAPL': 612.78, 'IBM': 205.55}
>>> tech_names = {'AAPL', 'IBM', 'HPQ', 'MSFT'}
>>> p2 = {k:v for k, v in prices.items() if k in tech_names }
>>> p2
{'HPQ': 37.2, 'AAPL': 612.78, 'IBM': 205.55}
>>> p1 = dict((k, v)for k, v in prices.items() if v > 200)
>>> p1
{'AAPL': 612.78, 'IBM': 205.55}
```
## 映射名称到序列的元素
* 为下标访问的序列构造名字，通过名字来访问该元素,使用命名元组提高代码的可读性
```python
>>> from collections import namedtuple                                     
>>> sub = namedtuple('sub', ['name', 'age'])
>>> c.name
'bob'
>>> c.age
'11'
```
* 可以使用命名元祖替代字典，命名元组不能直接赋值,可以使用_replace方法替换并重新生成一个命名元组
```python
>>> sub._replace(age='12')
info(name='bob', age='12')
```
* _replace() 方法还有一个很有用的特性就是当你的命名元组拥有可选或者缺失字段时候,它是一个非常方便的填充数据的方法。
```python
>>> info = namedtuple('info', ['name','age', 'other'])
>>> info_1 = info(None, None, None)
>>> def replace(s):
...     return info_1._replace(**s)
>>> 
>>> replace({'name':'bb', 'age':12, 'other':'111'})
info(name='bb', age=12, other='111')
```
## 转换并同时计算数据
* 结合数据计算与转换,使用一个生成器表达式参数
```python
>>> s = sum((x * x for x in nums)) #平方和
>>> s
55
```
* 一些使用test
```python
import os
files = os.listdir('.')
#any 任意一个为true就成立
if any(name.endswith('.py') for name in files):
    print('There be python')
else:
    print ('Sorry, no python')

s = ('ACME', 50, 123.45)
print (','.join(str(x) for x in s)) 
test = [ 
    {'name':'GOOG', 'shares':11},
    {'name':'GOOG', 'shares':75},
    {'name':'GOOG', 'shares':11}
]
#使用生成器作为函数参数可以不用重复使用括号
print (min(s['shares'] for s in test))
print (min((s['shares'] for s in test)))
#可以加入Key 配合匿名函数使用
print (min(test, key = lambda s:s['shares']))
```
## 多个映射合并为单个映射
* 检查a,b 字典，a如果没有去b中找
* ChainMap 会重新建立映射,重新定义常见的字典操作来进行操作
* 如果有重复的映射，只会使用第一个出现的映射
* 修改映射的值只会作用于第一个映射结构
```python
>>> a = {'x':1, 'z':3}
>>> b = {'y':2, 'z':4}
>>> from collections import ChainMap
>>> c = ChainMap(a, b)
>>> c
ChainMap({'x': 1, 'z': 3}, {'y': 2, 'z': 4})
>>> print (c['x'])
1
>>> print (c['y'])
2
>>> print (c['z'])
3
>>> list(c.keys())
['x', 'z', 'y']
>>> list(c.values())
[1, 3, 2]
>>> c['z'] = 10
>>> c
ChainMap({'x': 1, 'z': 10}, {'y': 2, 'z': 4})
```
* 可以使用dict 的update方法单独构造一个新字典,但如果存在相同的key新生成的字典只会有一个key,而且对原始数据进行修改的,不会反应到新生成的字典上,使用ChainMap就可以实现
```python
>>> a = {'x':1, 'z':3}
>>> b = {'y':2, 'z':4}
>>> maerged = ChainMap(a, b)
>>> maerged
ChainMap({'x': 1, 'z': 3}, {'y': 2, 'z': 4})
>>> maerged['x']
1
>>> a['x'] = 999
>>> maerged['x']
999
```







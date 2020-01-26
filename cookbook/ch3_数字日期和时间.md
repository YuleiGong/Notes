# 数字日期和时间

## 数字的四舍五入
* 简单的四舍五入使用round()
* 当一个值刚好在两个边界的中间的时候， round 函数返回离它最近的偶数。 也就是说，对1.5或者2.5的舍入运算都会得到2。
* 传给 round() 函数的 ndigits 参数可以是负数，这种情况下， 舍入运算会作用在十位、百位、千位等上面。比如
```python
>>> round(1.23, 1)
1.2
>>> round(1.27, 1)
1.3
>>> round(1.25361,3)
1.254
>>> a = 1627731
#个位
>>> round(a, -1)
1627730
#十位
>>> round(a, -2)
1627700
#百位
>>> round(a, -3)
1628000
>>>
```
## 精确的浮点数运算使用decimal
```python
>>> from decimal import Decimal
>>> a = Decimal('4.2')
>>> b = Decimal('2.1')
>>> a + b
Decimal('6.3')
>>> print(a + b)
6.3
>>> (a + b) == Decimal('6.3')
```
## 复数运算
```python
>>> a = complex(2, 4)
>>> b = 3 - 5j
>>> a
(2+4j)
>>> b
(3-5j)
#对应的实部、虚部和共轭复数可以很容易的获取
>>> a.real
2.0
>>> a.imag
4.0
>>> a.conjugate()
(2-4j)
```
* 如果要执行其他的复数函数比如正弦、余弦或平方根，使用 cmath 模块
```python
>>> import cmath
>>> cmath.sin(a)
(24.83130584894638-11.356612711218174j)
>>> cmath.cos(a)
(-11.36423470640106-24.814651485634187j)
>>> cmath.exp(a)
(-4.829809383269385-5.5920560936409816j)
>>>
```
## 分数运算
```
>>> from fractions import Fraction
>>> a = Fraction(5, 4)
>>> b = Fraction(7, 16)
>>> print(a + b)
27/16
>>> print(a * b)
35/64
```
## 大型数组运算
```python
>>> import numpy as np
>>> x = [1, 2, 3, 4]
>>> y = [5, 6, 7, 8]
>>> x * 2
[1, 2, 3, 4, 1, 2, 3, 4]
>>> x + 10
Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
    TypeError: can only concatenate list (not "int") to list
>>> x + y
[1, 2, 3, 4, 5, 6, 7, 8]
>>> import numpy as np
>>> ax = np.array([1, 2, 3, 4])
>>> ay = np.array([5, 6, 7, 8])
>>> ax * 2
array([2, 4, 6, 8])
>>> ax + 10
array([11, 12, 13, 14])
>>> ax + ay
array([ 6, 8, 10, 12])
>>> ax * ay
array([ 5, 12, 21, 32])
```
## 随机选择
* 从一个序列中随机抽取若干元素
```python
>>> import random
>>> values = [1, 2, 3, 4, 5, 6]
>>> random.choice(values)
2
>>> random.choice(values)
3
```
* 提取出N个不同元素的样本用来做进一步的操作
```python
>>> random.sample(values, 2)
[6, 2]
```
* 如果你仅仅只是想打乱序列中元素的顺序，可以使用 random.shuffle()
```python
>>> random.shuffle(values)
>>> values
[2, 4, 6, 5, 3, 1]
```
* 生成随机整数，请使用 random.randint()
```python
>>> random.randint(0,10)
2
>>> random.randint(0,10)
5
```
* 生成0到1范围内均匀分布的浮点数
```
>>> random.random()
0.9406677561675867
```
## 基本的日期与时间转换
* 为了执行不同时间单位的转换和计算，请使用 datetime 模块。
```python
>>> from datetime import timedelta
>>> from datetime import datetime
>>> a = timedelta(days=2, hours=6)
>>> b = timedelta(hours=4.5)
>>> c = a + b
>>> c.days
2
>>> c.seconds
37800
>>> c.seconds / 3600
10.5
>>> a = datetime(2012, 9, 23)
>>> print(a + timedelta(days=10))
2012-10-03 00:00:00
>>> b = datetime(2012, 12, 21)
>>> d = b - a
>>> d.days
89
>>> now = datetime.today()
>>> print(now)
2012-12-21 14:54:43.094063
>>> print(now + timedelta(minutes=10))
2012-12-21 15:04:43.094063
>>>
```
## 字符串转换为日期
```
>>> from datetime import datetime
>>> text = '2012-09-20'
>>> y = datetime.strptime(text, '%Y-%m-%d')
```




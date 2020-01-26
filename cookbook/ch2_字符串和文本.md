# 字符串和文本
## 针对任意多的分隔符拆分字符串
* 字符串中由多种不同的分隔符,也包括空格
```python
>>> import re
>>> line = 'asdf fjdk; afed, fjek,asdf, foo'
#[]匹配其中的任意一个字符,*(匹配前面的子表达式0次或1次)
>>> re.split(r'[;,\s]*', line)
```
## 字符串开头或结尾文本匹配
* 使用endswith和startswith做匹配
```python
>>> filename = 'spam.txt'
>>> filename.endswith('.txt')
True
>>> filename.startswith('file:')
False
>>> url = 'http://www.python.org'
>>> url.startswith('http:')
True
>>>
```
* 对多个选项做匹配,只需要将匹配条件放在tuple中即可, 必须使用tuple
* any 表示有任意一个成立即返回true
```python 
>>> import os
>>> filenames = os.listdir('.')
>>> filenames
[ 'Makefile', 'foo.c', 'bar.py', 'spam.c', 'spam.h' ]
>>> [name for name in filenames if name.endswith(('.c', '.h')) ]
['foo.c', 'spam.c', 'spam.h'
#可以用于检查某个文件夹下是否存在某种类型的文件
>>> any(name.endswith('.py') for name in filenames)
True
>>>
```
* 也可以使用正则表达式实现
```python
>>> import re
>>> url = 'http://www.python.org'
>>> re.match('http:|https:|ftp:', url)
<_sre.SRE_Match object at 0x101253098>
>>>
```
## 用Shell通配符匹配字符串
* 使用 Unix Shell 中常用的通配符(比如 *.py , Dat[0-9]*.csv 等)去匹配文本字符串
* 使用 fnmatch 模块提供了两个函数—— fnmatch() 和 fnmatchcase() 实现通配符匹配
```
>>> from fnmatch import fnmatch, fnmatchcase
>>> fnmatch('foo.txt', '*.txt')
True
>>> fnmatch('foo.txt', '?oo.txt')
True
>>> fnmatch('Dat45.csv', 'Dat[0-9]*')
True
>>> names = ['Dat1.csv', 'Dat2.csv', 'config.ini', 'foo.py']
>>> [name for name in names if fnmatch(name, 'Dat*.csv')]
['Dat1.csv', 'Dat2.csv']
>>>
```
* 可以使用fnmatchcase()来实现大小写的完全匹配
```python
>>> fnmatchcase('foo.txt', '*.TXT')
False
>>>
```
## 字符串匹配和搜索
* 简单的匹配使用find, endswith, startswith实现即可
* 复杂的匹配需要正则表达式
* 正则匹配核心步骤就是先使用 re.compile() 编译正则表达式字符串， 然后使用 match() , findall() 或者 finditer() 等方法。
```python
>>> text1 = '11/27/2012'
>>> import re
>>> if re.match(r'\d+/\d+/\d+', text1):
... print('yes')
... else:
... print('no')
yes
```
* 使用编译模式,用同一个正则规则去做多次匹配
```python
>>> datepat = re.compile(r'\d+/\d+/\d+')
>>> if datepat.match(text1):
... print('yes')
... else:
... print('no')
...
yes
```
* match()总是从字符串开始去匹配，一旦匹配到就结束。如果你想查找字符串任意部分的模式出现位置， 使用 findall() 方法去代替。
```python
>>> text = 'Today is 11/27/2012. PyCon starts 3/13/2013.'
>>> datepat.findall(text)
['11/27/2012', '3/13/2013']
>>>
```
* 使用括号去捕获分组
```python
datepat = re.compile(r'(\d+)/(\d+)/(\d+)')
>>> m = datepat.match('11/27/2012')
>>> m
<_sre.SRE_Match object at 0x1005d2750>
>>> # Extract the contents of each group
>>> m.group(0)
'11/27/2012'
>>> m.group(1)
'11'
>>> m.group(2)
'27'
>>> m.group(3)
'2012'
>>> m.groups()
('11', '27', '2012')
>>> month, day, year = m.groups()
>>> text
'Today is 11/27/2012. PyCon starts 3/13/2013.'
#在正则中使用了捕获分组以后,就会在findall中匹配到数据时按照规则组成元组
>>> datepat.findall(text)
[('11', '27', '2012'), ('3', '13', '2013')]
>>> for month, day, year in datepat.findall(text):
... print('{}-{}-{}'.format(year, month, day))
```
* 如果想使用精确匹配,在正则表达式中使用$结尾
```python
>>> datepat = re.compile(r'(\d+)/(\d+)/(\d+)$')
>>> datepat.match('11/27/2012abcdef')
>>> datepat.match('11/27/2012')
<_sre.SRE_Match object at 0x1005d2750>
```
## 字符串搜索和替换
* 简单的替换使用replace即可
* 对于复杂的替换,使用正则和sub()函数即可
* sub() 函数中的第一个参数正则，第二个参数是替换模式。反斜杠数字比如\3指捕获到的分组，可以任意改变顺序。
```python
>>> import re
>>> datepat = re.compile(r'(\d+)/(\d+)/(\d+)')
>>> datepat.sub(r'\3-\1-\2', text)
```
## 字符串忽略大小写的搜索替换
*  re 模块的时候给这些操作提供 re.IGNORECASE
```python
>>> text = 'UPPER PYTHON, lower python, Mixed Python'
>>> re.findall('python', text, flags=re.IGNORECASE)
['PYTHON', 'python', 'Python']
>>> re.sub('python', 'snake', text, flags=re.IGNORECASE)
'UPPER snake, lower snake, Mixed snake'
```
## 最短匹配模式(?非贪婪模式)
* __*__ 号的贪婪的会尽可能多的匹配
```python 
>>> str_pat = re.compile(r'\"(.*)\"')
>>> text1 = 'Computer says "no."'
>>> str_pat.findall(text1)
['no.']
>>> text2 = 'Computer says "no." Phone says "yes."'
>>> str_pat.findall(text2)
['no." Phone says "yes.']
>>>
```
* 使用非贪婪模式
```python
>>> str_pat = re.compile(r'\"(.*?)\"')
>>> str_pat.findall(text2)
['no.', 'yes.']
>>>
```
## 多行匹配模式
* re.compile() 函数接受一个标志参数叫 re.DOTALL ，在这里非常有用。 它可以让正则表达式中的点(.)匹配包括换行符在内的任意字符
```python
>>> text2 = '''/* this is a
... multiline comment */
... '''
>>> comment = re.compile(r'/\*(.*?)\*/', re.DOTALL)
>>> comment.findall(text2)
[' this is a\n multiline comment ']
```
## 删除字符串中不需要的字符
* strip() 方法能用于删除开始或结尾的字符。 lstrip() 和 rstrip() 分别从左和从右执行删除操作。 默认情况下，这些方法会去除空白字符，但是你也可以指定其他字符
```python
>>> s = ' hello world \n'
>>> s.strip()
'hello world'
>>> s.lstrip()
'hello world \n'
>>> s.rstrip()
' hello world'
>>> t = '-----hello====='
>>> t.lstrip('-')
'hello====='
>>> t.strip('-=')
'hello'
```
* 在删除操作中strip()不会对中间的文本产生影响
```python
>>> s = ' hello     world \n'
>>> s = s.strip()
>>> s
'hello     world'
>>>
```
* 借助replace或者sub可疑实现空格的全部删除
```python
>>> s.replace(' ', '')
'helloworld'
>>> import re
>>> re.sub('\s+', ' ', s)
'hello world'
>>>
```
* 通常情况下你想将字符串 strip 操作和其他迭代操作相结合，比如从文件中读取多行数据。在这里，表达式 lines = (line.strip() for line in f) 执行数据转换操作。 这种方式非常高效，因为它不需要预先读取所有数据放到一个临时的列表中去。 它仅仅只是创建一个生成器，并且每次返回行之前会先执行 strip 操作
```python
def open_file(filename):
    with open(filename) as f:
        lines = (line.strip() for line in f)
        for line in lines:
            print (line)
```
## 字符串对齐
* 函数 format() 可以用来很容易的对齐字符串
```python
>>> text = 'Hello World'
#右对齐20宽度
>>> format(text, '>20')
'         Hello World'
#左对齐20宽度
>>> format(text, '<20')
'Hello World         '
#中间对齐20宽度
>>> format(text, '^20')
'    Hello World     '
#指定填充字符串
>>> format(text, '=>20')
'=========Hello World'
>>> format(text, '*^20')
'****Hello World*****'
#格式化多个值
>>> '{:>10} {:>10}'.format('Hello', 'World')
'     Hello      World'
#格式化数字
>>> x = 1.2345
>>> format(x, '>10')
'    1.2345'
#保留2位小数
>>> format(x, '^10.2f')
'   1.23   '
```
## 合并拼接字符串
* 如果你想要合并的字符串是在一个序列或者 iterable 中，那么最快的方式就是使用 join() 方法。
```python
>>> parts = ['Is', 'Chicago', 'Not', 'Chicago?']
>>> ' '.join(parts)
'Is Chicago Not Chicago?'
>>> ','.join(parts)
'Is,Chicago,Not,Chicago?'
>>> ''.join(parts)
'IsChicagoNotChicago?'
>>> data = ['ACME', 50, 91.1]
#有非str类型的值,使用生成器表达式
>>> ','.join(str(d) for d in data)
'ACME,50,91.1'
```
* 准备编写构建大量小字符串的输出代码， 你最好考虑下使用生成器函数，利用yield语句产生输出片段
```python
def sample():
    yield 'Is'
    yield 'Chicago'
    yield 'Not'
    yield 'Chicago?'
if __name__ == '__main__':
    text = ''.join(sample())
    for part in sample():
        f.write(part)
```
## 字符串中插入变量
* 使用format方法
```python
>>> s = '{name} has {n} messages.'
>>> s.format(name='Guido', n=37)
```



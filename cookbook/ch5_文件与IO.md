# 读写文本数据
## 读写文本数据
* 读写各种不同编码的文本数据
* 使用带有 rt 模式的 open() 函数读取文本文件
* 写入一个文本文件，使用带有 wt 模式的 open() 函数
* 已存在文件中追加内容，使用模式为 at 的 open() 函数
* 文件的读写操作使用默认的系统编码，在读的时候最好转换为utf-8
* 查看系统的默认编码

```python
>>> import sys
>>> sys.getdefaultencoding()
'utf-8'
```
* 尽量使用with语句读取文本

```pyhton
with open('somefile.txt', 'rt', newline='') as f:
    pass
```
* 如果知道文件的编码格式,可以以指定的编码读取

```python
with open('somefile.txt', 'rt', encoding='latin-1') as f:
    pass
```
## 打印输出至文件中
* 在 print() 函数中指定 file 关键字参数

```python
def print_out():
    with open('test.txt', 'wt') as f:
        print("Hello World", file=f)
```

## 使用其他分隔符或行终止符打印
* 可以使用在 print() 函数中使用 sep 和 end 关键字参数

```python
>>> print ('aaa', 90, 51, sep=',', end='!!!')
aaa,90,51!!!
#在输出中禁止换行
>>> for i in range(5):
...     print (i, end=' ')
0 1 2 3 4 
```
## 读写字节数据
* 使用模式为 rb 或 wb 的 open() 函数来读取或写入二进制数据
* 字节字符串和文本字符串在迭代时返回值是不同的

```python
>>> t = "hello world"
>>> t[0]
'h'
>>> t = b"hello world"
>>> t[0]
104
```
* 从二进制模式的文件中读取或者写入数据，需要相应的解码和编码

```python
def wr_encode():
    with open('test.bin', 'wb') as f:
        f.write('hello world'.encode('utf-8'))
def rd_decode():
    with open('test.bin', 'rb') as f:
        data = f.read(16)
        text = data.decode('utf-8')
```

## 防止写的文件被覆盖
* 在 open() 函数中使用 x 模式来代替 w 模式来防止文件被重复写入

```python
>>> with open('test.txt', 'wt') as f:
...     f.write('Hello world')
>>> with open('test.txt', 'xt') as f:
...     f.write('Hello world')
Traceback (most recent call last):
  File "<input>", line 1, in <module>
      with open('test.txt', 'xt') as f:
      FileExistsError: [Errno 17] File exists: 'test.txt'
#也可以在写前面检测文件是否存在
>>> import os
>>> os.path.exists('test.txt')
True
```
## 固定大小的文件迭代
* 在一个固定长度记录或者数据块的集合上迭代，而不是在一个文件中一行一行的迭代
* 不断的产生固定长度的可迭代数据块,直到结束。如果迭代的字节大小不是整数倍，最后一次迭代数据字节数会小
* b''代表达到文件结尾的返回值
* 必须以rb 二进制模式打开

```python
from functools import partial
RECORD_SIZE = 32
def size_rb():
    with open('somefile.data', 'rb') as f:
        records = iter(partial(f.read, RECORD_SIZE), b'')
        for r in records:
            print (r)
```
## 读取二进制数据到可变缓冲区中
* 直接读取二进制数据到一个可变缓冲区中，而不需要做任何的中间复制操作
* 文件对象的 readinto() 方法能被用来为预先分配内存的数组填充数据
* 返回实际读取的字节数

```python
import os
def read_into_buffer(filename):
    #设定缓冲区大小
    buf = bytearray(os.path.getsize(filename))
    with open(filename, 'rb') as f:
        f.readinto(buf)
    return buf

if __name__ == '__main__':
    with open('sample.bin', 'wb') as f:
        f.write(b'Hello World')
    #读取数据到缓冲区
    buf = read_into_buffer('sample.bin')
    with open('newsample.bin', 'wb') as f:
        f.write(buf)
```
## 内存映射的二进制文件
* 使用 mmap 模块来内存映射文件
* mmap() 返回的 mmap 对象同样也可以作为一个上下文管理器来使用,这时候底层的文件会被自动关闭。
* 默认情况下， memeory_map() 函数打开的文件同时支持读和写操作。 任何的修改内容都会复制回原来的文件中。 如果需要只读的访问模式，可以给参数 access 赋值为 mmap.ACCESS_READ
* 如果你想在本地修改数据，但是又不想将修改写回到原始文件中，可以使用 mmap.ACCESS_COPY
* 多个Python解释器内存映射同一个文件，得到的 mmap 对象能够被用来在解释器直接交换数据。 也就是说，所有解释器都能同时读写数据，并且其中一个解释器所做的修改会自动呈现在其他解释器中

```python
import os
import mmap

def memory_map(filename, access=mmap.ACCESS_WRITE):
    size = os.path.getsize(filename)
    fd = os.open(filename, os.O_RDWR)
    return mmap.mmap(fd, size, access=access)

if __name__ == '__main__':
    size = 1000000
    with open('data', 'wb') as f:
        #设置文件的起始位置
        f.seek(size-1)
        f.write(b'\x00')

    with memory_map('data') as m:
        print (len(m))
        m[0:11] = b'hello world'
        print (m[0:11])
```

## 文件路径名的操作
* 使用os.path

```python
>>> import os
>>> path = '/Users/beazley/Data/data.csv'
>>> os.path.basename(path)
'data.csv'
>>> os.path.dirname(path)
'/Users/beazley/Data'
>>> os.path.join('tmp', 'data', os.path.basename(path))
'tmp/data/data.csv'
>>> path = '~/Data/data.csv'
>>> os.path.expanduser(path)
'/Users/gongyulei/Data/data.csv'
>>> os.path.splitext(path)
('~/Data/data', '.csv')
```
## 测试文件是否存在

```python
>>> os.path.exists('/etc/passwd')
True
>>> os.path.exists('/tmp/spam')
False
>>> os.path.exists('/tmp/')
True
>>> os.path.isfile('/etc/passwd')
True
>>> os.path.isdir('/etc/passwd')
False
#软连接测试
>>> os.path.islink('/usr/local/bin/python3')
True
#文件的真实路径
>>> os.path.realpath('/usr/local/bin/python3')
'/usr/local/Cellar/python3/3.6.1/Frameworks/Python.framework/Versions/3.6/bin/python3.6'
#获取大小
>>> os.path.getsize('/etc/passwd')
5925
```
## 获取文件夹中的文件列表

```python
#结果会返回目录中所有文件列表
>>> os.listdir('.')
['ch5_10.py', 'ch5_2.py', 'ch5_4.py', 'ch5_8.py', 'ch5_9.py', 'data', 'newsample.bin', '
sample.bin', 'somefile.data', 'test.bin', 'test.gz', 'test.txt']
#执行某种过滤
>>> name = [name for name in os.listdir() if os.path.isfile(os.path.join('.', name))]
>>> name
['ch5_10.py', 'ch5_2.py', 'ch5_4.py', 'ch5_8.py', 'ch5_9.py', 'data', 'newsample.bin', '
sample.bin', 'somefile.data', 'test.bin', 'test.gz', 'test.txt']
>>> name = [name for name in os.listdir() if os.path.isdir(os.path.join('.', name))]
>>> name
['test']
>>> pyfils = [name for name in os.listdir() if name.endswith('.py')]
>>> pyfils
['ch5_10.py', 'ch5_2.py', 'ch5_4.py', 'ch5_8.py', 'ch5_9.py']
```
* 收集文件的其他数据

```python
import os
import glob

def main():
    pyfiles = glob.glob('*.py')
    print (pyfiles)
    name_sz_date = [(name, os.path.getsize(name), os.path.getmtime(name))\
                   for name in pyfiles]
    for name, size, mtime in name_sz_date:
        print (name, size, mtime)

if __name__ == '__main__':
    main()
```

## 增加或改变已打开文件的编码
* 给一个以二进制模式打开的文件添加Unicode编码/解码方式,借助io.TextIOWrapper()
* io.TextIOWrapper()是一个编码和解码Unicode的文本处理层

```python
import urllib.request
import io

def encode_text():
    u = urllib.request.urlopen('http://www.python.org')
    f = io.TextIOWrapper(u, encoding='utf-8')
    #编码成utf-8
    text = f.read()
    print (text)

if __name__ == '__main__':
    encode_text()
```
## 创建临时文件和文件夹
* 在程序执行时创建一个临时文件或目录，并希望使用完之后可以自动销毁掉
* w+t'为文本模式,delete=False，表示临时文件不会被删除

```python
from tempfile import NamedTemporaryFile, TemporaryDirectory

def create_tmp():
    #'w+t'为文本模式
    with NamedTemporaryFile('w+t', delete=False) as f:
        print ('filename is :', f.name)
        f.write('Hello World\n')
        f.write('Testing\n')
        f.seek(0)
        data = f.read()

    with TemporaryDirectory() as dirname:
        print ('dirname is :', dirname)

if __name__ == '__main__':
    create_tmp()

```
* 可以自己定制自己的临时文件规则

```python
>>> from tempfile import NamedTemporaryFile
>>> f = NamedTemporaryFile(prefix='mytemp', suffix='.txt', dir='/tmp')
>>> f.name
'/tmp/mytempurb_uyz4.txt'
>>> 
```
## 序列化python对象
* 使用dump写入，load解包。写入和解包的数据类型可以是类，List等

```python
>>> import pickle
>>> f = open('test.txt', 'wb')
>>> pickle.dump([1, 2, 3], f)
>>> pickle.dump({'aaa', '111'}, f)
>>> f = open('test.txt', 'rb')
>>> pickle.load(f)
[1, 2, 3]
>>> pickle.load(f)
{'111', 'aaa'}
```
* 某些对象是无法进行pickle的,比如打开的网络连接，线程。class 可以通过__getstate__和__setstats__解决



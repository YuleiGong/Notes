# 数据编码与处理
## 读写csv文件
* 通过元组,命名元组,字典方式来读取csv

```pyhton
import csv
from collections import namedtuple
def main1():
    with open('stocks.csv') as f:
        f_csv = csv.reader(f)
        #跳过表头
        headers = next(f_csv)
        for row in f_csv:
            print (row)

def main2():
    with open('stocks.csv') as f:
        f_csv = csv.reader(f)
        headings = next(f_csv)
        Row = namedtuple('Row', headings)
        for r in f_csv:
            row = Row(*r)
            print (row)

def main3():
    with open('stocks.csv') as f:
        f_csv = csv.DictReader(f)
        for r in f_csv:
            print (r)
```
* 在使用命名元组时,需要处理表头头非法字符的情况比如'-',使用正则进行替换

```python
def main5():
    import re
    with open('stocks.csv') as f:
        f_csv = csv.reader(f)
        headers = [re.sub(r'[^a-zA-Z]', '_', h) for h in next(f_csv)]
        Row = namedtuple('Row', headers)
        for r in f_csv:
            row = Row(*r)
            print (row)
```

* 写入csv, 分别以列表和字典的形式写入

```python
def write_1():
    headers = ['Symbol','Price','Date','Time','Change','Volume']
    rows = [
        ('AA', 39.48, '6/11/2007', '9:36am', -0.18, 181800),
        ('AIG', 71.38, '6/11/2007', '9:36am', -0.15, 195500),
        ('AXP', 62.58, '6/11/2007', '9:36am', -0.46, 935000),
    ]
    with open('stocks.csv', 'w') as f:
        f_csv = csv.writer(f)
        f_csv.writerow(headers)
        for row in rows:
            f_csv.writerow(row)

def write_2():
    headers = ['Symbol', 'Price', 'Date', 'Time', 'Change', 'Volume']
    rows = [
        {'Symbol':'AA', 'Price':39.48, 'Date':'6/11/2007','Time':'9:36am', 'Change':-0.18, 'Volume':181800},
        {'Symbol':'AIG', 'Price': 71.38, 'Date':'6/11/2007','Time':'9:36am', 'Change':-0.15, 'Volume': 195500},
        {'Symbol':'AXP', 'Price': 62.58, 'Date':'6/11/2007','Time':'9:36am', 'Change':-0.46, 'Volume': 935000},
    ]
    with open('stocks.csv', 'w') as f:
        f_csv = csv.DictWriter(f, headers)
        f_csv.writeheader()
        for row in rows:
            f_csv.writerow(row)

```
* 改变编码的读取规则,例如以tab键分隔的csv

```python
def main4():
    with open('stocks.csv') as f:
        #tab 分隔
        f_csv = csv.reader(f, delimiter='\t')
        for row in f_csv:
            print (row)

```
* csv是不会对数据进行额外处理,需要自行处理

```python
def convert_1():
    col_types = [str, float, str, str, float, int]
    with open('stocks.csv') as f:
        f_csv = csv.reader(f)
        headers = next(f_csv)
        for row in f_csv:
            row = tuple(convert(value) for convert, value in zip(col_types, row))
            print (row)

def convert_2():
    field_types = [
        ('Price', float),
        ('Change', float),
        ('Volume', int)
    ]
    with open('stocks.csv') as f:
        for row in csv.DictReader(f):
            #每行逐次扫描,转换数据更新到字典中
            row.update((key, conversion(row[key])) for key, conversion in field_types)
            print (row)
```
## 读写Josn数据
* json 解码会解码出字典或者列表,在loads时传递object_pairs_hook或object_hook参数,可以解码成需要的对象
* 分别将json字符串解码成OrderedDict和JSONObject对象

```python
s= '{"name": "ACME", "shares": 50, "price": 490.1}'
>>> from collections import OrderedDict
>>> json.loads(s, object_pairs_hook=OrderedDict)
OrderedDict([('name', 'ACME'), ('shares', 50), ('price', 490.1)])
>>> class JSONObject:
...     def __init__(self, d):
...         self.__dict__ = d
>>> data = json.loads(s, object_hook=JSONObject)
>>> data.name
'ACME'
>>> 
```
* 是编码json输出变得好看

```python
>>> print (json.dumps({'a':1}, indent=4))
{
    "a": 1
}
```
* 序列化类实例,通过函数将类实例转化为字典

```python
import json
class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

def serialize_instance(obj):
    d = {'__classname__':type(obj).__name__}
    d.update(vars(obj))
    return d

if __name__ == '__main__':
    p = Point(2, 3)
    s = json.dumps(p, default=serialize_instance)
    print (s)
```
## 解析简单的xml
* parse()将整个xml文档解析为文档对象,就可以利用find查询特定的信息

```python
>>> from urllib.request import urlopen
>>> from xml.etree.ElementTree import parse
>>> u = urlopen('http://planet.python.org/rss20.xml')
>>> doc = parse(u)
>>> doc
<xml.etree.ElementTree.ElementTree object at 0x10dcd42e8>
>>> e = doc.find('channel/title')
>>> e.tag
'title'
>>> e.text
'Planet Python'
```

## 将字典转化为xml
* 如果需要保持dict元素的顺序,需要使用OrdereDict对象

```python
from xml.etree.ElementTree import Element, tostring
def dict_to_xml(tag, d):
    #创建最外层的节点
    elem = Element(tag)
    for key, val in d.items():
        child = Element(key)
        child.text = str(val)
        #最外层节点上添加内容
        elem.append(child)

    return elem

if __name__ == '__main__':
    s = {'name':'GOOG', 'shares':100, 'price':490.1}
    e = dict_to_xml('stock', s)
    #给原始添加属性
    e.set('_id', '1234')
    print (tostring(e))
#输出
b'<stock _id="1234"><name>GOOG</name><shares>100</shares><price>490.1</price></stock>'
```
* 使用字符串去构造xml

```python
def dict_to_xml_str(tag, d):
    parts = ['<{}>'.format(tag)]
    for key, val in d.items():
        parts.append('<{0}>{1}<0>'.format(key, val))
    parts.append('</{}>'.format(tag))
    return ''.join(parts)

if __name__ == '__main__':
    s = {'name':'GOOG', 'shares':100, 'price':490.1}
    e = dict_to_xml_str('stock', s)
    print (e)
#输出
<stock><name>GOOG<0><shares>100<0><price>490.1<0></stock>
```

## 编码和解码十六进制数
* 字节字符串和十六进制的编码或解码
* base64中的16进制转换只能操作大写形式

```python
>>> s = b'hello'
>>> import binascii
>>> h = binascii.b2a_hex(s)
>>> h
b'68656c6c6f'
>>> binascii.a2b_hex(h)
b'hello'
>>> import base64
>>> h = base64.b16encode(s)
>>> h
b'68656C6C6F'
>>> base64.b16decode(h)
b'hello'
#编码为Unicode
>>> h = h.decode('ascii')
>>> h
'68656C6C6F'
```
## encode(编码) decode(解码) base64

```python
>>> import base64
>>> s = b'hello'
>>> a = base64.b64encode(s)
>>> a
b'aGVsbG8='
>>> base64.b64decode(a)
b'hello'
#解码为unicode
>>> base64.b64decode(a).decode('ascii')
'hello'
```
## 读写二进制数组数据
* 写入元组到二进制文本中
* 使用struct编码写入

```python
from struct import Struct
def write_records(records, format, f):
    record_struct = Struct(format)
    for r in records:
        f.write(record_struct.pack(*r))

if __name__ == '__main__':
    #write
    records = [
        (1, 2.3, 4.5),
        (6, 7.8, 9.0),
        (12, 13.4, 56.7),
    ]
    with open('data.b', 'wb') as f:
        #小端存储 int double double
        write_records(records, '<idd', f)
```
* 以增量形式读取二进制数据

```python
def read_records(format, f):
    record_struct = Struct(format)
    #没有形参的lambda,不断的生成size大小的迭代器,直到b''为止
    chunks = iter(lambda: f.read(record_struct.size), b'')
    #生成器可以迭代三次,每次20字节
    return (record_struct.unpack(chunk) for chunk in chunks)

if __name__ == '__main__':
    #read
    with open('data.b', 'rb') as f:
        for r in read_records('<idd', f):
            print (r)
```
* 全量读取二进制文件

```python
def unpack_records(format, data):
    '''
    设置好解包步长,一次解包返回迭代器
    '''
    record_struct = Struct(format)
    return (record_struct.unpack_from(data, offset) \
           for offset in range(0, len(data), record_struct.size))

if __name__ == '__main__':
    #read
    with open('data.b', 'rb') as f:
        data = f.read()
        for rec in unpack_records('<idd', data):
            print (rec)
```





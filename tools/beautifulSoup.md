# beautifulSoup
## 基本使用

```
#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-03-30 15:50:04
from __future__ import unicode_literals
from __future__ import absolute_import

from bs4 import BeautifulSoup

html = '''
<html><head><title>The Dormouse's story</title></head>
<body>
<p class="title"><b>The Dormouse's story</b></p>

<p class="story">Once upon a time there were three little sisters; and their names were
<a href="http://example.com/elsie" class="sister" id="link1">Elsie</a>,
<a href="http://example.com/lacie" class="sister" id="link2">Lacie</a> and
<a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>;
and they lived at the bottom of a well.</p>
<p class="story">...</p>
'''

if __name__ == '__main__':
    soup = BeautifulSoup(html,'lxml')
    for link in soup.find_all('a'):
        print (link.get('href'))
    print (soup.get_text())
    #print(soup.prettify())
    #print(soup.title)
    #print(soup.title.name)
    #print(soup.title.string)
    #print(soup.title.parent.name)
    #print(soup.p)
    #print(soup.p["class"])
    #print(soup.a)
    #print(soup.find_all('a'))
    #print(soup.find(id='link3'))
```
## contents 获取子标签
* p.contents 会将标签的子标签存入列表中,p.children同理，只不过会返回生成器

```
from __future__ import unicode_literals
from __future__ import absolute_import

html = """
<html>
    <head>
        <title>The Dormouse's story</title>
    </head>
    <body>
        <p class="story">
            Once upon a time there were three little sisters; and their names were
            <a href="http://example.com/elsie" class="sister" id="link1">
                <span>Elsie</span>
            </a>
            <a href="http://example.com/lacie" class="sister" id="link2">Lacie</a>
            and
            <a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>
            and they lived at the bottom of a well.
        </p>
        <p class="story">...</p>
"""

from bs4 import BeautifulSoup

soup = BeautifulSoup(html,'lxml')
#获取P标签下的虽有子标签，存入列表
print(soup.p.contents)
print (soup.p.children)
for i, child in enumerate(soup.p.children):
    print (i, child)
```
## 标准选择器
* find_all,可以根据标签名，属性，内容查找

```
html='''
<div class="panel">
    <div class="panel-heading">
        <h4>Hello</h4>
    </div>
    <div class="panel-body">
        <ul class="list" id="list-1">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
            <li class="element">Jay</li>
        </ul>
        <ul class="list list-small" id="list-2">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
        </ul>
    </div>
</div>
'''
from bs4 import BeautifulSoup
soup = BeautifulSoup(html, 'lxml')
print(soup.find_all('ul'))
#再次遍历查找
for ul in soup.find_all('ul'):
    print (ul.find_all('li'))
print(soup.find_all('ul')[0])
```
* attrs attrs可以传入字典的方式来查找标签

```
html='''
<div class="panel">
    <div class="panel-heading">
        <h4>Hello</h4>
    </div>
    <div class="panel-body">
        <ul class="list" id="list-1" name="elements">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
            <li class="element">Jay</li>
        </ul>
        <ul class="list list-small" id="list-2">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
        </ul>
    </div>
</div>
'''
from bs4 import BeautifulSoup
soup = BeautifulSoup(html, 'lxml')
print(soup.find_all(attrs={'id': 'list-1'}))
print(soup.find_all(attrs={'name': 'elements'}))
```

## css选择器
* .表示class #表示id

```
html='''
<div class="panel">
    <div class="panel-heading">
        <h4>Hello</h4>
    </div>
    <div class="panel-body">
        <ul class="list" id="list-1">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
            <li class="element">Jay</li>
        </ul>
        <ul class="list list-small" id="list-2">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
        </ul>
    </div>
</div>
'''
from bs4 import BeautifulSoup
soup = BeautifulSoup(html, 'lxml')
print(soup.select('.panel .panel-heading'))
print(soup.select('ul li'))
print(soup.select('#list-2 .element'))
```

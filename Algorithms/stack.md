
# 栈

* __栈__ 是有序集合,添加和移除操作总发生在同一端。即__顶端__
* 栈中的元素离底端越近,代表在栈中的时间越长。最新添加的元素将最先被移除。这种排序原则称为 LIFO(后进先出last-in-first-out)。
* 最近添加的元素靠近顶端，旧元素则靠近底端。元素的插入顺序和移除顺序正好相反。
* 浏览器的返回按钮设计就是一个栈,当前流浪的网页放置于栈顶端,最早浏览的网页位于底端。

* python实现,使用列表来模拟栈,使用列表的__底端__作为栈的__顶端__

```python
class Stack:
    """
    栈的python实现,当栈增长(push),新的元素会被添加到列表的尾部。pop同样会修改这一端
    """
    def __init__(self):
        self.items = []

    def isEmpty(self):
        return self.items == []

    def push(self,item):
        self.items.append(item)

    def pop(self):
        return self.items.pop()

    def peek(self):
        """
        永远返回列表的最后一位,模拟出栈的后进先出
        """
        return self.items[len(self.items) - 1]

    def size(self):
        return len(self.items)
out:
>>> from stack import Stack
>>> s = Stack()
>>> s.isEmpty()
True
>>> s.size()
0
>>> s.push(1)
>>> s.push(2)
>>> s.push(3)
>>> s.pop()
3
>>> s.pop()
2
>>> s.peek()
1
```

* python实现,使用列表来模拟栈,使用列表的__顶端__作为栈的__顶端__

```python
class Stack:
    """
    栈的python实现
    """
    def __init__(self):
        self.items = []

    def isEmpty(self):
        return self.items == []

    def push(self,item):
        """
        永远在顶端添加数据
        Args:
            items:入栈的数据
        """
        self.items.insert(0,item)

    def pop(self):
        return self.items.pop(0)

    def peek(self):
        return self.items[0]

    def size(self):
        return len(self.items)
```


## 通过栈来匹配括号
* 我们需要通过使用栈来匹配 诸如 ``` (()()()()) ``` 这样的括号。
* 如果匹配到左括号,使用push入栈,否则使用pop出栈。遍历完成后,如果栈为空,证明括号都是匹配的。


```python
def par_checker(symbolstring):
    """
    判断左右括号是否匹配
    Args:
        symbolstring:包含左右括号的列表
    """
    s = Stack()
    balanced = True
    index = 0

    while index < len(symbolstring) and balanced:
        symbol = symbolstring[index]
        if symbol == "(":
            s.push(symbol)
        else:
            if s.isEmpty():
                balanced = False
            else:
                s.pop()
        index = index + 1

    if balanced and s.isEmpty():
        return True
    else:
        return False

if __name__ == '__main__':
    symbolstring = list(("((()))"))
    print (par_checker(symbolstring))
```

## 通过栈来匹配符号
* 我们需要通过使用栈来匹配 诸如:```([][])``` 这样的符号,python中的字典,列表,都需要这样的匹配
* 如果匹配到左符号,使用push入栈,匹配到右符号时,需要检测类型,如果类型一致使用pop出栈。遍历完成后,如果栈为空,证明符号都是匹配的。

```python
def par_checker(symbolstring):
    s = Stack()
    balanced = True
    index = 0

    while index < len(symbolstring) and balanced:
        symbol = symbolstring[index]
        if symbol in "([{":
            s.push(symbol)
        else:
            top = s.pop()
            if not matches(top, symbol):
                balanced = False

        index = index + 1

    if balanced and s.isEmpty():
        return True
    else:
        return False

def matches(open,close):
    opens = "([{"
    closers = ")]}"


    return opens.index(open) == closers.index(close)

if __name__ == '__main__':
    symbolstring = "({[)})"
    print (par_checker(symbolstring))
```
## 十进制数转换成二进制数
* 使用一种除以2的算法,循环不停的将10进制数除以2,余数就是二进制结构。计算出的第一个余数是最后一位,利用栈的__后进先出__策略来存储

```python
from stack import Stack

def divide_by2(dec_number):
    remstack = Stack()

    while dec_number > 0:
        rem = dec_number % 2 #余数部分
        remstack.push(rem)
        dec_number = dec_number // 2 #整数部分

    bin_string = ""
    while not remstack.isEmpty():
        bin_string = bin_string + str(remstack.pop())

    return bin_string


if __name__ == '__main__':
    print (divide_by2(233))
```

* 十进制转换为任意进制数,只需修改除以2位除以一个基数。当基数超过10,需要一套新的数字来代表余数(余数本身就是十进制)


```python
def divide_converter(dec_number, base):
    """
    十进制数转换成任意进制数字
    Args:
        dec_number:十进制数
        base:转换基数
    Returns:
        返回base进制的数
    """
    digits = "0123456789ABCDEF"

    remstack = Stack()

    while dec_number > 0:
        rem = dec_number % base #余数部分
        remstack.push(rem)
        dec_number = dec_number // base #整数部分

    new_string = ""
    while not remstack.isEmpty():
        new_string = new_string + digits[remstack.pop()]

    return new_string
```

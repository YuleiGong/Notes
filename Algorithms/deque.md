# 双端队列
* 双端队列是与队列类似的有序集合。它有一前、一后两端，元素在其中保持自己的位置。与队列不同的是，双端队列对在哪一端添加和移除元素没有任何限制。新元素既可以被添加到前端,也可以被添加到后端。同理，已有的元素也能从任意一端移除。
<a href="https://sm.ms/image/25iP8SrQbvKJIoL" target="_blank"><img src="https://i.loli.net/2020/01/24/25iP8SrQbvKJIoL.png" ></a>

* python 实现队列前端进行移入和移除的时间复杂度是O(1),队列后端移入移除的时间复杂度是O(n)

```python
class Deque:
    def __init__(self):
        self.items = []

    def isEmpty(self):
        return self.items  == []

    def addFront(self,item):
        self.items.append(item)

    def addRear(self,item):
        self.items.insert(0,item)

    def removeFront(self):
        return self.items.pop()

    def removeRear(self):
        return self.items.pop(0)

    def size(self):
        return len(self.items)
```

## 回文检测器 
* __回文__ 指的是一个字符串从前往后,或者从后往前都是一样的字符串。例如toot,radar
* 双端队列可以从前后两端移除元素,如果前端移除元素与后端移出元素,一致(元素个数为偶数),或者最后只剩下一个元素(奇数)

```python
from deque import Deque

def palchecker(aString):
    """
    回文检测
    Args:
        aString:待检测的字符串
    """
    chardeque = Deque()
    for ch in aString:
        chardeque.addRear(ch)

    stillEqual = True

    while chardeque.size() > 1 and stillEqual:
        first = chardeque.removeFront()
        last = chardeque.removeRear()
        if first != last:
            stillEqual = False

    return stillEqual

if __name__ == '__main__':
    print (palchecker('dsfkdekjke'))
    print (palchecker('toot'))
```


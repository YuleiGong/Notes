# 列表
* __列表__ 是元素的集合，其中每一个元素都有一个相对于其他元素的位置。更具体地说，这种列表称为无序列表。

## 无序列表

* python实现,为了实现无序列表,我们需要构建__链表__,__无序列表__需要维护元素之间的相对位置,但并不需要在连续的内存空间中维护这些位置信息。如果可以为每一个元素维护下一个元素的位置,那么这些元素的相对位置就能通过下一个元素来指定。

<a href="https://sm.ms/image/Ssm5KXjMuex1V6f" target="_blank"><img src="https://i.loli.net/2020/01/26/Ssm5KXjMuex1V6f.png" ></a>

* 无序列表的抽象表示
<a href="https://sm.ms/image/KY5SrF6VvufQpxA" target="_blank"><img src="https://i.loli.net/2020/01/26/KY5SrF6VvufQpxA.png" ></a>
* __Node__ 代表列表的基本数据结构,next存储了下一个节点数据的__引用__
* __UnorderedList__ 代表无序列表数据结构,head代表列表头。
* 删除元素后,需要将上一个元素的引用和下一个元素的引用相链接。

```python
class Node:
    """
    构建列表的基本数据结构
    """

    def __init__(self, initdata):
        self.data = initdata
        self.next = None

    def getData(self):
        return self.data

    def getNext(self):
        return self.next

    def setData(self, newdata):
        self.data = newdata

    def setNext(self,newnext):
        self.next = newnext

class UnorderedList:
    """
    列表对象
    """

    def __init__(self):
        self.head = None

    def isEmpty(self):
        return self.head == None

    def add(self, item):
        """
        列表中添加元素
        Args:
            item: 添加的元素
        """
        temp = Node(item)
        temp.setNext(self.head)
        self.head = temp

    def length(self):
        """
        列表的长度统计
        """
        current = self.head
        count = 0
        while current != None:
            count = count + 1
            current = current.getNext()
        return count

    def search(self,item):
        """
        查找元素是
        Args:
            item:需要查找的元素
        Retuens:
            False-未找到 True-找到
        """
        current = self.head
        found = False
        while current != None and not found:
            if current.getData() == item:
                found = True
            else:
                current = current.getNext()

        return found

    def remove(self, item):
        """
        移除某个节点,将目标节点前面的节点和后面的节点做连接
        Args:
            item:需要查找的元素
        """
        current = self.head
        previous = None
        found = False
        while not found and current != None:
            if current.getData() == item:
                found = True
            else:
                previous = current
                current = current.getNext()
        if previous == None and found == True:
            #第一个元素就找到
            self.head = current.getNext()
        elif found == True:
            previous.setNext(current.getNext())
        else:
            pass

```

## 有序列表
* 有序列表中,元素的相对位置取决于他们的基本特征。通常以升序或降序排列
* 有序列表和无序列表的不同:
    * search 只需要找到比目标元素大的节点就可以终止
    * add 需要根据顺序来插入,不能简单的将一个列表放入头部

```python
class Node:
    """
    构建列表的基本数据结构
    """

    def __init__(self, initdata):
        self.data = initdata
        self.next = None

    def getData(self):
        return self.data

    def getNext(self):
        return self.next

    def setData(self, newdata):
        self.data = newdata

    def setNext(self,newnext):
        self.next = newnext



class OrderedList:
    def __init__(self):
        self.head = None

    def search(self,item):
        current = self.head
        found = False
        stop = False
        while current != None and not found and not stop:
            if current.getData() == item:
                found = True
            else:
                if current.getData() > item:
                    stop = True
                else:
                    current = current.getNext()

        return found

    def add(self, item):
        current = self.head
        previous = None
        stop = False

        while current != None and not stop:
            if current.getData() > item:
                stop = True
            else:
                if current.getData() > item:
                    stop = True
                else:
                    previous = current
                    current = current.getNext()

        temp = Node(item)
        if previous == None:
            temp.setNext(self.head)
            self.head = temp
        else:
            temp.setNext(current)
            previous.setNext(temp)
```

## 时间复杂度分析
* n个节点的链表为例， isEmpty 方法的时间复杂度是O(1) ，这是因为它只需要执行一步操作，即检查head引用是否为None。
* length 方法则总是需要执行n步操作，这是因为只有完全遍历整个列表才能知道究竟有多少个元素。因此，length 方法的时间复杂度是O(n)。
* 向无序列表中添加元素是O(1),因为只是简单地将新节点放在链表的第一个位置。
* 有序列表的 search、remove 以及 add 都需要进行遍历操作。尽管它们平均都只需要遍历一半的节点,但是这些方法的时间复杂度都是O(n)。在最坏的情况下它们都需要遍历所有节点。

 

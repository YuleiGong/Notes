# 搜索

* __搜索__ 是指从元素集合中找到某个特定元素的算法过程,搜索函数通常放回True或False,用来表示元素是否存在

## 顺序搜索
* 存储于列表等集合中的数据项彼此存在__线性或顺序__的关系，每个数据项的位置与其他数据项相关。在Python列表中.数据项的位置就是它的下标。因为下标是有序的，所以能够顺序访问,由此可以进行顺序搜索
* __无序列表__ 的顺序搜索 python实现
    * 因为是无序列表的搜索,至多需要搜索N次,时间复杂度O(n)

```python
def seqentialSearch(alist, item):
    """
    无序列表的顺序搜索
    Args:
        alist:待搜索的列表
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    pos = 0
    found = False

    while pos < len(alist) and not found:
        if alist[pos] == item:
            found = True
        else:
            pos = pos + 1

    return found
```
* __有序列表__  的顺序搜索 python实现
    * 对比无序列表的搜索,只有在不存在目标元素时,普遍的情况可能只需要遍历n/2次,所以最坏情况的时间复杂的还是O(n)


```python
def orderedSequentialSearch(alist, item):
    """
    顺序列表的顺序搜索
    Args:
        alist:待搜索的列表
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    pos = 0
    found = False
    stop = False
    while pos < len(alist) and not found and not stop:
        if alist[pos] == item:
            found = True
        else:
            if alist[pos] > item:
                stop = True
            else:
                pos = pos + 1

    return found
```
## 二分搜索
* __二分搜索__

```python
def binarySearch(alist, item):
    """
    有序列表的二分搜索
    Args:
    Args:
        alist:待搜索的列表(有序列表)
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    first = 0
    last = len(alist) - 1
    found = False

    while first <= last and not found:
        midpoint = (first + last) // 2
        if alist[midpoint] == item:
            found = True
        else:
            if item < alist[midpoint]:
                last = midpoint - 1
            else:
                first = midpoint + 1
    return found

```
* __二分搜索__  的递归版本

```python
def binarySearch1(alist, item):
    """
    有序列表的二分搜索(递归版本)
    Args:
    Args:
        alist:待搜索的列表(有序列表)
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    if len(alist) == 0:
        return False
    else:
        midpoint = len(alist) // 2
        if alist[midpoint] == item:
            return True
        else:
            if item < alist[midpoint]:
                return binarySearch1(alist[:midpoint], item)
            else:
                return binarySearch1(alist[midpoint+1:], item)
```
* 二分搜索中,第一次比较剩余n/2,第二次n/4,类推n/2^i(i代表次数),i=logn,由此可得,二分搜索算法的时间复杂的是对数O(logn)。
* 尽管二分搜索通常优先于顺序搜索,但当数据过大,额外的排序开销是否值得。对于大型列表而言,排序也会造成巨大的计算成本。

## 散列
* __散列表__ 是元素的集合,其中的元素以一种便于查找的方式存储。散列表中的每个位置被称为__槽__ 。
<a href="https://sm.ms/image/IKcph5LEkqw2f3i" target="_blank"><img src="https://i.loli.net/2020/01/28/IKcph5LEkqw2f3i.png" ></a>
* __散列函数__ 将散列表中的元素与其所属位置对于起来。对散列表中的任一元素,散列函数返回一个介于 0 和 m – 1 之间的整数。
    * 假设在一个固定大小为11的散列表中,假设散列函数是 item%11(散列表大小)
    <a href="https://sm.ms/image/Ez1LDQXgHYFTkCx" target="_blank"><img src="https://i.loli.net/2020/01/28/Ez1LDQXgHYFTkCx.png" ></a>
    * 该散列表的分布,其中有6个槽被占用,占有率λ=元素个数/散列表大小:
    <a href="https://sm.ms/image/5prt7q1bmjlXdN9" target="_blank"><img src="https://i.loli.net/2020/01/28/5prt7q1bmjlXdN9.png" ></a>
* 搜索元素时,只需要根据散列函数的值,就可以找到对应槽中的元素。时间复杂的为O(1)
* 如果元素为44 散列值为0,就会和77冲突,这也叫 __碰撞__。
* __散列函数__不发生碰撞的称为 __完美散列函数__,可以通过增加散列大小来降低碰撞几率,如果元素过少,这是可行的,如果元素过多,会极大的浪费空间。我们的目标是创建这样一个散列函数:冲突数最少,计算方便,元素均匀分布于散列表中,有多种常见的方法来扩展取余函数:
    * 常见的散列函数取余方法有: __折叠法__ __平法取中法__

    * 一种字符串散列函数
    ```python
    def hash(astring, tablesize):
        """
        一种字符串散列函数
        Args:
            astring:需要散列的字符串
            tablesize:散列表大小
        Returns:
            散列值
        """
        sum = 0
        for pos in range(len(astring)):
            ch = ord(astring[pos])
            sum = sum + ch*(pos+1)
        return sum % tablesize
    ```
* __处理散列冲突__:
    * 一种方法是在散列表中找到另一个空槽,用于放置引起冲突的元素。简单的做法是从起初的散列值开始,顺序遍历散列表,直到找到一个空槽。为了遍历散列表,可能需要往回检查第一个槽。这个过程被称为__开放定址法__。 它尝试在散列表中寻找下一个空槽或地址。由于是逐个 访问槽，因此这个做法被称作 __线性探测__。线性探测可能会使散列表中元素出现 __聚集__ 现象,如果一个元素的散列值重复,可能会在该元素周围的槽上补充元素。
    * 一种方法是 __再散列__ 泛指在发生冲突后寻找另一个槽的过程。可以将再散列函数定义为 rehash(pos) = (pos + skip)%sizeoftable。注意，__跨步(skip)__的大小要能保证表中所有的槽最终都被访问到，否则就会浪费槽资源。
    * 一种方法是 __平方探测__ 。它不采用固定的跨步大小，而是通过再散列函数递增散列 值。如果第一个冲突散列值是h，后续的散列值就是h+1^2、h+2^2、h+3^、h+4^2...
    * 一种方法是 __链接法__ 。链接法允许在散列上链接一个列表引用,元素全部存入其中。计算得到散列值后,需要再次搜索列表得到元素
    <a href="https://sm.ms/image/bqmwj4E1rAkf3t7" target="_blank"><img src="https://i.loli.net/2020/01/28/bqmwj4E1rAkf3t7.png" ></a>

* __字典__ 的散列表实现 python

```python
class HashTable:
    """
    构造一个散列表
    """
    def __init__(self):
        self.size = 11
        self.slots = [None] * self.size
        self.data = [None] * self.size

    def put(self, key, data):
        """
        散列表中存入值
        Args:
            key:key
            data:data
        """
        hashvalue = self.hashfunction(key, len(self.slots))

        if self.slots[hashvalue] == None:
            self.slots[hashvalue] = key
            self.data[hashvalue] = data
        else:
            if self.slots[hashvalue] == key:
                self.data[hashvalue] = data
            else:
                nextsolt = self.rehash(hashvalue,len(self.slots))
                while self.slots[nextsolt] != None and \
                      self.slots[nextsolt] != key:
                    nextsolt = self.rehash(nextsolt,len(self.slots))
                if self.slots[nextsolt] == None:
                    self.slots[nextsolt] = key
                    self.data[nextsolt] = data
                else:
                    self.data[nextsolt] = data

    def hashfunction(self,key,size):
        """
        计算散列值
        """
        return key % size

    def rehash(self, oldhash, size):
        return (oldhash + 1) % size

    def get(self, key):
        startslot = self.hashfunction(key, len(self.slots))
        
        data = None
        stop = False
        found = False
        position = startslot
        while self.slots[position] != None and \
              not found and not stop:
            if self.slots[position] == key:
                found = True
                data = self.data[position]
            else:
                position = self.rehash(position, len(self.slots))
                if position == startslot:
                    stop = True

        return data

    def __getitem__(self, key):
        return self.get(key)
    
    def __setitem__(self, key, data):
        return self.put(key, data)
```

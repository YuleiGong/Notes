# 排序
* __排序__ 是指将集合中的元素按某种顺序排列的过程

## 冒泡排序
* __冒泡排序__ 多次遍历列表,它比较相邻的元素,将不符合顺序的交换,每轮遍历都将下一个最大值放到正确的位置。
* 第一轮比较n-1对,第二轮比较n-2对,每次都找到一个最大的元素

<a href="https://sm.ms/image/lu3JU4M9BHNvpea" target="_blank"><img src="https://i.loli.net/2020/01/28/lu3JU4M9BHNvpea.png" ></a>
* __冒泡排序__ python 实现,冒泡排序的时间复杂度为O(n^2),速度很慢

```python
def bubbleSort(alist):
    """
    列表的冒泡排序
    Args:
        alist:待排序的列表
    """
    #倒序遍历
    for passnum in range(len(alist)-1, 0,-1):
        for i in range(passnum):
            if alist[i] > alist[i+1]:
                alist[i],alist[i+1] = alist[i+1],alist[i]
```
* __短冒泡__ python实现,在每一轮的遍历中,如果该轮没有做冒泡交换,排序结束

```python
def shortBubbleSort(alist):
    """
    短冒泡排序,每一轮遍历,一旦发现不需要冒泡,退出
    Args:
        alist:待排序的列表
    """
    exchange = True
    passnum = len(alist) - 1
    while passnum > 0 and exchange:
        exchange = False
        for i in range(passnum):
            if alist[i] > alist[i+1]:
                exchange = True
                alist[i],alist[i+1] = alist[i+1],alist[i]

        passnum = passnum - 1
```
## 选择排序
* __选择排序__ 在冒泡的基础上做了改进,每次遍历只做一次交换。第一次遍历寻找最大值,之后遍历完放到正确的位置上。第二次遍历寻找第二大,同样放到正确的位置上,以此类推。
<a href="https://sm.ms/image/MV7jhd1iIgJYqts" target="_blank"><img src="https://i.loli.net/2020/01/28/MV7jhd1iIgJYqts.png" ></a>

```python
def selectionSort(alist):
    """
    选择排序
    Args:
        alist:需要排序的元素列表
    """
    #逆序遍历
    for fillslot in range(len(alist)-1,0,-1):
        positionOfMax = 0
        for location in range(1, fillslot+1):
            if alist[location] > alist[positionOfMax]:
                positionOfMax = location
        temp = alist[fillslot]
        alist[fillslot],alist[positionOfMax] = alist[positionOfMax],alist[fillslot]
```
* 虽然 __选择排序__ 和 __冒泡排序__ 轮次相同,时间复杂度也是O(n^2),但选择排序减小了每一轮的交互次数

## 插入排序
* __插入排序__ 的时间复杂度也是O(n^2)。在列表的较低的一端维护一个有序的子列表,并逐个将新元素插入 __子列表__

<a href="https://sm.ms/image/sUDwVBj5f8cYay3" target="_blank"><img src="https://i.loli.net/2020/01/28/sUDwVBj5f8cYay3.png" alt="插入排序.png"></a>

* 在 __子列表__ 中将比它大的元素右移;如果遇到一个比它小的元素,就可以插入当前元素,该论排序结束。

<a href="https://sm.ms/image/QEoKaC9VtbB1gPA" target="_blank"><img src="https://i.loli.net/2020/01/28/QEoKaC9VtbB1gPA.png" ></a>

```python
def insertionSort(alist):
    """
    插入排序
    Args:
        alist:待排序的无序列表
    """


    for index in range(1, len(alist)):
        currentvalue = alist[index]
        position = index

        while position > 0 and alist[position - 1] > currentvalue:
            alist[position] = alist[position-1]
            position = position - 1

        alist[position] = currentvalue
```
## 希尔排序
* __希尔排序__ 也称为 __递减增量排序__ ,他对插入排序做了改进,并对每一个子列表应用插入排序。如何切分列表是希尔排序的关键——并不是连续切分,而是使用增量i(有时称作步长)选取所有间隔为i的元素组成子列表。

* 使用步长3作为列表切分间隔,在排序完成后,使用一次插入排序,完成排序。
<a href="https://sm.ms/image/UekzNIDmQqg6cpH" target="_blank"><img src="https://i.loli.net/2020/01/29/UekzNIDmQqg6cpH.png" ></a>

* __希尔排序__ python实现,第一次排序n/2个子序列,第一次排列n/4个子序列...以此类推。直到最后进行一次间隔为1的插入排序

```python
def shellSort(alist):
    """
    希尔排序
    Args:
        alist:待排列的元素列表
    """
    sublistcount = len(alist) // 2

    while sublistcount > 0:
        for startposition in range(sublistcount):
            gapInsertionSort(alist,startposition,sublistcount)

        print ("After increments of size",sublistcount,"The list is,",alist)

        sublistcount = sublistcount // 2

def gapInsertionSort(alist,start, gap):
    """
    按照step插入排序
    Args:
        alist:待排列的元素列表
        start:排序间隔
        gap:元素总数
    """
    for i in range(start+gap,len(alist),gap):
        currentvalue = alist[i]
        position = i
        while position >= gap and \
              alist[position-gap] > currentvalue:
            alist[position] = alist[position-gap]
            position = position - gap

        alist[position] = currentvalue
```
* __希尔排序__ 是每一轮遍历的结果都生成了更有序的列表,这是的最后一步的插入排序非常高效

## 归并排序
* __归并排序__ 属于递归算法,每次将一个列表一分为二,直到不能在一分为二,这一步骤称为 __拆分__ 。最后进行 __归并__ 这一操作,将较小的有序列表归并为一个有序列表。

<a href="https://sm.ms/image/ukSGYEKdnPjOApl" target="_blank"><img src="https://i.loli.net/2020/01/29/ukSGYEKdnPjOApl.png" ></a>

## 快速排序 TODO

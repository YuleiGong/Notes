# 队列
* __队列__是有序集合,添加发生在__尾部__,移除操作则发生在__头部__。新元素从尾部进入队列,然后一直向前移动到头部,直到成为下一个被移除的元素。
* 新添加的元素必须在队列的尾部等待。在队列中时间最长的元素则排在最前面，这种原则被称为__FIFO(fitst-in-first-out)先进先出__

* python实现

```python
class Queue:
    def __init__(self):
        self.items = []

    def isEmpty(self):
        return self.items == []

    def enqueue(self,item):
        """
        队列的尾部添加一个元素,这里使用list的第一个元素代表尾部,因为list的pop是从最好有个元素开始。所以用list[0]代表队列的尾部
        Args:
            item:队列尾部添加的元素

        """
        self.items.insert(0,item)

    def dequeue(self):
        return self.items.pop()

    def size(self):
        return len(self.items)

if __name__ == '__main__':
    q = Queue()
    q.enqueue(1)
    q.enqueue(2)
    q.enqueue(3)
    print (q.dequeue())
    print (q.size())
```

## 传土豆模拟
* 考虑这样一个 儿童游戏:传土豆。在这个游戏中，孩子们围成一圈，并依次尽可能快地传递一个土豆，如图所示。在某个时刻，大家停止传递，此时手里有土豆的孩子就得退出游戏。重复上述过程，直到只剩下一个孩子。
<a href="https://sm.ms/image/DhKFytECbaPmORN" target="_blank"><img src="https://i.loli.net/2020/01/14/DhKFytECbaPmORN.png" ></a>
* 使用队列来模拟一个环,每运动一次,头部的数据移动到尾部,在某一时刻(num=出列入列次数),位于队列头部的数据出局。
<a href="https://sm.ms/image/Azaxm5gwPst7lkD" target="_blank"><img src="https://i.loli.net/2020/01/14/Azaxm5gwPst7lkD.png" ></a>


```python
def hot_potato(namelist, num):
    simqueue = Queue()
    for name in namelist:
        simqueue.enqueue(name)

    while simqueue.size() > 1:
        for i in range(num):
            simqueue.enqueue(simqueue.dequeue())
        #num次循环完成后,出局的数据
        simqueue.dequeue()

    return simqueue.dequeue()

if __name__ == '__main__':
    print (hot_potato(["Bill", "David", "Susan", "Jane", "Kent", "Brad"],7))
```

## 打印任务模拟
* 向一台共享打印机发送任务,任务被存储在队列中,并且按照__先进先出__的顺序进行打印。打印机每分钟可以打印的页数是固定的。
* 
* 打印机对象 __tick__ 方法模拟打印机运行。

```python
class Printer:
    """
    打印机
    """
    def __init__(self, ppm):
        """
        初始化打印速度
        Args:
            ppm:每分钟打印多少页
        """
        self.pagerate = ppm
        self.currentTask = None
        self.timeRemaining = 0

    def tick(self):
        if self.currentTask != None:
            self.timeRemaining = self.timeRemaining - 1
            if self.timeRemaining <= 0:
               self.currentTask = None

    def busy(self):
        """
        任务繁忙检测
        Returns:
            打印机繁忙,返回True
            打印机空闲,返回False
        """
        if self.currentTask != None:
            return True
        else:
            return False

    def startNext(self, newtask):
        self.currentTask = newtask
        #打印花费的时间 单位s
        self.timeRemaining = newtask.getPages() * 60 /self.pagerate
```
* 任务对象

```python
class Task:
    def __init__(self, time):
        self.timestamp = time
        self.pages = random.randrange(1,21)

    def getStamp(self):
        return self.timestamp

    def getPages(self):
        return self.pages

    def waitTime(self,currenttime):
        """
        任务等待的时间
        Args:
            currenttime:当前时间
        """
        return currenttime - self.timestamp
```

* 模拟打印任务:3600秒内,任务需要消耗的平均等待时间


```python
def simulation(numSeconds,pagesPerMinute):
    """
    打印任务模拟
    Args:
        numSeconds:numSeconds 时间内模拟打印
        pagesPerMinute:打印速度 pages/min
    """
    labprinter = Printer(pagesPerMinute)
    printQueue = Queue()
    waitingtimes = []

    for currentSecond in range(numSeconds):
        if newPrintTask():
            task = Task(currentSecond)
            printQueue.enqueue(task)
        #打印机空闲且队列不为空
        if (not labprinter.busy()) and (not printQueue.isEmpty()):
            #下一个任务
            nexttask = printQueue.dequeue()
            waitingtimes.append(nexttask.waitTime(currentSecond))
            labprinter.startNext(nexttask)
        labprinter.tick()

    averageWait = sum(waitingtimes)/len(waitingtimes)
    print ("Average wait %6.2f secs %3d tasks remaining."%(averageWait,printQueue.size()))

def newPrintTask():
    """
    平均每180秒一个随机任务
    """
    num = random.randrange(1,181)
    if num == 180:
        return True
    else:
        return False

if __name__ == '__main__':
    for i in range(10):
        simulation(3600,5)
```


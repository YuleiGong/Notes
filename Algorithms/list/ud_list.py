#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-26 15:26:34

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


if __name__ == '__main__':
    l = UnorderedList()
    l.add(1)
    l.add(2)
    l.add(3)
    l.add(4)
    l.add(5)
    print (l.length())
    print (l.search(3))
    l.remove(10)
    print (l.length())


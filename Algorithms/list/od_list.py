#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-26 16:46:36
from __future__ import unicode_literals
from __future__ import absolute_import

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

if __name__ == '__main__':
    l = OrderedList()
    l.add(1)
    l.add(3)
    l.add(5)
    l.add(6)
    l.add(2)
    print (l.search(3))










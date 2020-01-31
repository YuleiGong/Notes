#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-31 14:04:00

class BinHeap:
    """
    二叉堆的实现
    Attributes:
        heapList:初始化二叉堆列表
        currentSize:二叉堆的大小
    """

    def __init__(self):
        self.heapList = [0]
        self.currentSize = 0

    def buildHeap(self,alist):
        """
        通过列表构建二叉堆
        """
        i = len(alist) // 2
        self.currentSize = len(alist)
        self.heapList = [0] + alist[:]
        while (i>0):
            self.percDown(i)
            i = i - 1


    def insert(self,k):
        """
        二叉堆中插入新元素,列表中插入元素,可以保证完全数的性质,但破坏了堆的结构。
        需要逐次比较与父元素的大小,移动元素。
        Args:
            k:插入的元素
        """
        self.heapList.append(k)
        self.currentSize = self.currentSize + 1
        self.percUp(self.currentSize)

    def percUp(self, i):
        while i // 2 > 0:
            if self.heapList[i] < self.heapList[i // 2]:
                tmp = self.heapList[i // 2]
                self.heapList[i // 2] = self.heapList[i]
                self.heapList[i] = tmp
            i = i // 2

    def delMin(self):
        """
        移除堆的最小元素:
        列表的根节点是最小元素,可以直接移除。移除后,需要保证二叉堆的结构性和有序性
        结构性:取出列表中的最后一个元素,将其移动到根节点的位置。
        有序性:逐次比较子节点,移动位置。保证堆的有序性
        Returns:
            retval:返回最小元素
        """
        retval = self.heapList[1]
        self.heapList[1] = self.heapList[self.currentSize]
        self.currentSize = self.currentSize - 1
        self.heapList.pop()
        self.percDown(1)
        return retval

    def percDown(self,i):
        while (i * 2) <= self.currentSize:
            mc = self.minChild(i)
            if self.heapList[i] > self.heapList[mc]:
                tmp = self.heapList[i]
                self.heapList[i] = self.heapList[mc]
                self.heapList[mc] = tmp
            i = mc


    def minChild(self, i):
        if i * 2 + 1 > self.currentSize:
            return i * 2
        else:
            if self.heapList[i*2] < self.heapList[i*2+1]:
                return i * 2
            else:
                return i * 2 + 1

if __name__ == '__main__':
    bh = BinHeap()
    bh.insert(5)
    bh.insert(7)
    bh.insert(3)
    bh.insert(11)
    print (bh.delMin())

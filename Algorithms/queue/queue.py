#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-14 21:24:37

class Queue:
    def __init__(self):
        self.items = []

    def isEmpty(self):
        """
        队列为空,返回True
        """
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





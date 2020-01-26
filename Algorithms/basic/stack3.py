#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-13 23:40:52
from __future__ import unicode_literals
from __future__ import absolute_import

class MyStack:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.items = []
        

    def push(self, x: int) -> None:
        """
        Push element x onto stack.
        """
        self.items.append(x)

    def pop(self) -> int:
        """
        Removes the element on top of the stack and returns that element.
        """
        self.items.pop()
        

    def top(self) -> int:
        """
        Get the top element.
        """
        return self.items[len(self.items)-1]
        

    def empty(self) -> bool:
        """
        Returns whether the stack is empty.
        """
        return self.items == []
        


#Your MyStack object will be instantiated and called as such:
x = 1
obj = MyStack()
obj.push(x)
#param_2 = obj.pop()
param_3 = obj.top()
param_4 = obj.empty()

#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-01 15:59:54
from __future__ import unicode_literals
from __future__ import absolute_import

class Stack:
    """
    栈的python实现
    """
    def __init__(self):
        self.items = []

    def isEmpty(self):
        return self.items == []

    def push(self,item):
        """
        永远在顶端添加数据
        Args:
            items:入栈的数据
        """
        self.items.insert(0,item)

    def pop(self):
        return self.items.pop(0)

    def peek(self):
        return self.items[0]

    def size(self):
        return len(self.items)



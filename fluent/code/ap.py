#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-08 20:47:28
from __future__ import unicode_literals
from __future__ import absolute_import

""""
等差数列类实现
"""

class ArithmeticProgression:

    def __init__(self,begin,step,end=None):
        self.begin = begin
        self.step = step
        self.end = end

    def __iter__(self):
        result = type(self.begin + self.step)(self.begin)
        forever = self.end is None #True生成无穷序列
        index = 0
        while forever or result < self.end:
            yield result
            index += 1
            result = self.begin + self.step * index

def aritprog_gen(begin, step, end=None):
    result = type(begin + step)(begin)
    forever = end is None
    index = 0
    while forever or result < end:
        yield result
        index += 1
        result = begin + step * index

if __name__ == '__main__':
    pass





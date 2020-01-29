#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-28 20:26:52
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

if __name__ == '__main__':
    alist = [5,4,1,2,9]
    shortBubbleSort(alist)
    print (alist)


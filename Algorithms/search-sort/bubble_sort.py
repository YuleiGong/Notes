#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-28 20:26:52
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


if __name__ == '__main__':
    alist = [5,4,1,2,9,1]
    bubbleSort(alist)
    print (alist)


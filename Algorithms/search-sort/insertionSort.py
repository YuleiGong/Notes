#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-28 22:38:02

def insertionSort(alist):
    """
    插入排序
    Args:
        alist:待排序的无序列表
    """


    for index in range(1, len(alist)):
        currentvalue = alist[index]
        position = index

        while position > 0 and alist[position - 1] > currentvalue:
            alist[position] = alist[position-1]
            position = position - 1

        alist[position] = currentvalue

if __name__ == '__main__':
    alist = [54,26,93,17,77,31,44,55,20]
    insertionSort(alist)
    print (alist)

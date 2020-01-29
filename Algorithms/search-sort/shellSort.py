#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-29 14:31:59

def shellSort(alist):
    """
    希尔排序
    Args:
        alist:待排列的元素列表
    """
    sublistcount = len(alist) // 2

    while sublistcount > 0:
        for startposition in range(sublistcount):
            gapInsertionSort(alist,startposition,sublistcount)

        print ("After increments of size",sublistcount,"The list is,",alist)

        sublistcount = sublistcount // 2

def gapInsertionSort(alist,start, gap):
    """
    按照step插入排序
    Args:
        alist:待排列的元素列表
        start:排序间隔
        gap:元素总数
    """
    for i in range(start+gap,len(alist),gap):
        currentvalue = alist[i]
        position = i
        while position >= gap and \
              alist[position-gap] > currentvalue:
            alist[position] = alist[position-gap]
            position = position - gap

        alist[position] = currentvalue

if __name__ == '__main__':
    alist = [54, 26, 93, 17, 77, 31, 44, 55, 20]
    shellSort(alist)
    print (alist)


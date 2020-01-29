#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-28 21:30:17

def selectionSort(alist):
    """
    选择排序
    Args:
        alist:需要排序的元素列表
    """
    #逆序遍历
    for fillslot in range(len(alist)-1,0,-1):
        positionOfMax = 0
        for location in range(1, fillslot+1):
            if alist[location] > alist[positionOfMax]:
                positionOfMax = location
        temp = alist[fillslot]
        alist[fillslot],alist[positionOfMax] = alist[positionOfMax],alist[fillslot]


if __name__ == '__main__':
    alist = [5,4,6,9]
    selectionSort(alist)
    print (alist)




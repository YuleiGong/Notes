#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-28 12:34:57

def binarySearch(alist, item):
    """
    有序列表的二分搜索
    Args:
    Args:
        alist:待搜索的列表(有序列表)
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    first = 0
    last = len(alist) - 1
    found = False

    while first <= last and not found:
        midpoint = (first + last) // 2
        if alist[midpoint] == item:
            found = True
        else:
            if item < alist[midpoint]:
                last = midpoint - 1
            else:
                first = midpoint + 1
    return found

def binarySearch1(alist, item):
    """
    有序列表的二分搜索(递归版本)
    Args:
    Args:
        alist:待搜索的列表(有序列表)
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    if len(alist) == 0:
        return False
    else:
        midpoint = len(alist) // 2
        if alist[midpoint] == item:
            return True
        else:
            if item < alist[midpoint]:
                return binarySearch1(alist[:midpoint], item)
            else:
                return binarySearch1(alist[midpoint+1:], item)




if __name__ == '__main__':
    print (binarySearch([1,2,3,4,5],2))
    print (binarySearch1([1,2,3,4,5],2))





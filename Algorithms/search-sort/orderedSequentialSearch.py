#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-28 12:15:41
from __future__ import unicode_literals
from __future__ import absolute_import

def orderedSequentialSearch(alist, item):
    """
    顺序列表的顺序搜索
    Args:
        alist:待搜索的列表
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    pos = 0
    found = False
    stop = False
    while pos < len(alist) and not found and not stop:
        if alist[pos] == item:
            found = True
        else:
            if alist[pos] > item:
                stop = True
            else:
                pos = pos + 1

    return found

if __name__ == '__main__':
    print (orderedSequentialSearch([1,2,4,5],2))
    print (orderedSequentialSearch([1,2,4,5],3))


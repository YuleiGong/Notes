#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-27 22:31:30

def seqentialSearch(alist, item):
    """
    无序列表的顺序搜索
    Args:
        alist:待搜索的列表
        item:搜索的元素
    Returns
        found:True-元素存在 False-元素不存在
    """
    pos = 0
    found = False

    while pos < len(alist) and not found:
        if alist[pos] == item:
            found = True
        else:
            pos = pos + 1

    return found

if __name__ == '__main__':
    print (seqentialSearch([1,2,3,4],1))


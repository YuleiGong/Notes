#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-24 15:45:58

from deque import Deque

def palchecker(aString):
    """
    回文检测
    Args:
        aString:待检测的字符串
    """
    chardeque = Deque()
    for ch in aString:
        chardeque.addRear(ch)

    stillEqual = True

    while chardeque.size() > 1 and stillEqual:
        first = chardeque.removeFront()
        last = chardeque.removeRear()
        if first != last:
            stillEqual = False

    return stillEqual

if __name__ == '__main__':
    print (palchecker('dsfkdekjke'))
    print (palchecker('toot'))

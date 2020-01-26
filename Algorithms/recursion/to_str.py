#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-09 21:55:50
from __future__ import unicode_literals
from __future__ import absolute_import

"""
1.不断的做除法运算取余数,直到小于基数
2.通过查表法将余数对应的字符串取出来

"""


def to_str(n,base):
    converString = '0123456789ABCDEF'
    if n<base:
        return converString[n]
    else:
        return to_str(n/base,base) + converString[n%base]

if __name__ == '__main__':
    print (to_str(1,2))

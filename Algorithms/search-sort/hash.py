#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-28 15:31:01
from __future__ import unicode_literals
from __future__ import absolute_import

def hash(astring, tablesize):
    """
    一种字符串散列函数
    Args:
        astring:需要散列的字符串
        tablesize:散列表大小
    Returns:
        散列值
    """
    sum = 0
    for pos in range(len(astring)):
        ch = ord(astring[pos])
        sum = sum + ch*(pos+1)

    return sum % tablesize


if __name__ == '__main__':
    print (hash('cat',11))


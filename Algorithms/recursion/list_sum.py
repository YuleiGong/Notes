#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-09 21:31:50
from __future__ import unicode_literals
from __future__ import absolute_import

"""
递归的列表求和
[1,2,3,4,5]
"""
def list_sum(num_list):
    if len(num_list) == 1:
        return num_list[0]
    else:
        print (num_list)
        return num_list[0] + list_sum(num_list[1:])

if __name__ == '__main__':

    print (list_sum([1,3,5,7,9]))


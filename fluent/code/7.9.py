#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-02-17 16:20:35
from __future__ import unicode_literals
from __future__ import absolute_import

def make_averager():
    series = []
    def averager(new_value):
        series.append(new_value)
        total = sum(series)
        return total / len(series)
    return averager


if __name__ == '__main__':
    avg = make_averager()
    print (avg(10))
    print (avg(11))

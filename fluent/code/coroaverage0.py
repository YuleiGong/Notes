#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-10 09:08:24
from __future__ import unicode_literals
from __future__ import absolute_import
from coroutil import coroutine


@coroutine
def averager():
    total = 0
    count = 0
    average = None
    while True:
        term = yield average
        total += term
        count += 1
        average = total/count

if __name__ == '__main__':
    pass

#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-10 17:54:08
from __future__ import unicode_literals
from __future__ import absolute_import

from collections import namedtuple

Result = namedtuple('Result', 'count average')

def averager():
    total = 0.0
    count = 0
    average = None
    while True:
        term = yield
        if term is None:
            break
        total += term
        count += 1
        average = total/count
    return Result(count, average)

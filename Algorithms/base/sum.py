#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-01 12:44:05
from __future__ import unicode_literals
from __future__ import absolute_import

import time

def sumOfN(n):
    start = time.time()

    theSum = 0
    for i in range(1,n+1):
        theSum = theSum + i

    end = time.time()

    return theSum,end-start

def sumOfN3(n):
    start = time.time()
    _sum = (n*(n+1))/2
    end = time.time()

    return _sum,end-start


if __name__ == '__main__':
    for i in range(5):
        print ("Sum is %d required %10.7f seconds" % sumOfN3(100000))

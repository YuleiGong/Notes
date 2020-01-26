#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-09 17:35:28
from __future__ import unicode_literals
from __future__ import absolute_import

def simple_coroutine():
    print ('-> coroutine started')
    x = yield
    print ('-> coroutine received:', x)

def simple_coro2(a):
    """
    产出两个值的协程
    """
    print ('-> Started: a=',a)
    b = yield a
    print ('-> Received: b=',b)
    c = yield a + b
    print ('-> Received: c=',c)

if __name__ == '__main__':
    pass

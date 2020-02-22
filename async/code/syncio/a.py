#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-22 11:59:34
from __future__ import unicode_literals
from __future__ import absolute_import

import asyncio
import datetime
import time

def func1(end_time, loop):
    print ("func1 called")
    if (loop.time() + 1.0) < end_time:
        loop.call_later(1, func2, end_time, loop)
    else:
        loop.stop()

def func2(end_time, loop):
    print ("func2 called")
    if (loop.time() + 1.0) < end_time:
        loop.call_later(1, func3, end_time, loop)
    else:
        loop.stop()

def func3(end_time, loop):
    print ("func3 called")
    if (loop.time() + 1.0) < end_time:
        loop.call_later(1, func1, end_time, loop)
    else:
        loop.stop()




if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    end_loop = loop.time() + 9.0
    loop.call_soon(func1, end_loop, loop)
    loop.run_forever()
    loop.close()

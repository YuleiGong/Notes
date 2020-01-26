#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-10 16:26:37
from __future__ import unicode_literals
from __future__ import absolute_import

class DemoException(Exception):
    pass

def demo_exc_handling():
    print ('-> coroutine started')
    while True:
        try:
            x = yield
        except DemoException:
            print ('*** DemoException handled.Continuing...')
        else:
            print ('-> coroutine received:{!r}'.format(x))
    raise RuntimeError('This line should never run.')


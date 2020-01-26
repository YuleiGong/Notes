#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-09 16:03:11
from __future__ import unicode_literals
from __future__ import absolute_import

import contextlib

@contextlib.contextmanager
def looking_glass():
    import sys
    original_write = sys.stdout.write

    def reverse_write(text):
        original_write(text[::-1])

    sys.stdout.write = reverse_write
    msg = ''
    try:
        yield 'JABBERWOCKY'
    except ZeroDivisionError:
        msg = 'Please DO NOT divide by zero'
    finally:
        sys.stdout.write = original_write
        if msg:
            print (msg)

if __name__ == '__main__':
    pass

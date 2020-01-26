#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-09 09:25:27
from __future__ import unicode_literals
from __future__ import absolute_import

def chain(*iterbales):
    for i in iterbales:
        yield from i


if __name__ == '__main__':
    pass


#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-09 10:21:02
from __future__ import unicode_literals
from __future__ import absolute_import

import random

def d6():
    return random.randint(1,6) 

if __name__ == '__main__':
    d6_iter = iter(d6,1)
    for roll in d6_iter:
        print (roll)

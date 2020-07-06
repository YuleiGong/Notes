#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-04-13 22:06:45
from __future__ import unicode_literals
from __future__ import absolute_import

import bisect
import random

SIZE = 7

random.seed(1729)

my_list = []
for i in range(SIZE):
    new_item = random.randrange(SIZE*2)
    bisect.insort(my_list, new_item)
    print('%2d ->' % new_item, my_list)

#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-16 16:10:56
from __future__ import unicode_literals
from __future__ import absolute_import

import time

def countdown(n):
    while n > 0:
        print ("T-minus", n)
        n -= 1
        time.sleep(5)
from threading import Thread

t = Thread(target=countdown, args=(3,),daemon=True)
t.start()

if t.is_alive():
    print ('Still runing')
else:
    print ('Completed')

#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-17 10:40:53
from __future__ import unicode_literals
from __future__ import absolute_import

from threading import Thread, Event
import time


def countdown(n, started_evt):
    print ('countding starting')
    started_evt.set()
    while n > 0:
        print ('T-minus', n)
        n -= 1
        time.sleep(5)

started_evt = Event()

print ('Launching countdown')
t = Thread(target=countdown, args=(10, started_evt))
t.start()
started_evt.wait()

print ('countdown is runing')

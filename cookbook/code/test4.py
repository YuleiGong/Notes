#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-18 10:49:34
from __future__ import unicode_literals
from __future__ import absolute_import

import threading
import time

class PeriodicTime:
    def __init__(self, interval):
        self._interval = interval
        self._flag = 0
        self._cv = threading.Condition()

    def start(self):
        t = threading.Thread(target=self.run)
        t.daemon = True
        t.start()

    def run(self):
        while True:
            time.sleep(self._interval)
            with self._cv:
                self._flag ^= 1
                self._cv.notify_all()

    def wait_fo_tick(self):
        with self._cv:
            last_flag = self._flag
            while last_flag == self._flag:
                self._cv.wait()

def countdown(nticks):
    while nticks > 0:
        ptimer.wait_fo_tick()
        print('T-minus', nticks)
        nticks -= 1

def countup(last):
    n = 0
    while n < last:
        ptimer.wait_fo_tick()
        print ('Counting',n)
        n += 1


if __name__ == '__main__':
    ptimer = PeriodicTime(5)
    ptimer.start()
    threading.Thread(target=countdown, args=(10,)).start()
    threading.Thread(target=countup, args=(5,)).start()












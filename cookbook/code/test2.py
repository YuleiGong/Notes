#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-17 10:14:27
from __future__ import unicode_literals
from __future__ import absolute_import

import time
from threading import thread

class countdowntask:
    def __init__(self):
        self._runing = true

    def terminate(self):
        self._runing = false

    def run(self,n):
        while self._runing and n > 0:
            print ('t-minus', n)
            n -= 1
            time.sleep(5)

if __name__ == '__main__':
    c = countdowntask()
    t = thread(target=c.run, args=(10,))
    t.start()
    c.terminate()
    t.join()



#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 16:24:47
from __future__ import unicode_literals
from __future__ import absolute_import

import threading
import time
import random

semaphore = threading.Semaphore(0)

def consumer():
    print ("consumer is waiting")
    semaphore.acquire()#-1
    print ("Consumer notify: consumed item number {}".format(item))

def producer():
    global item
    time.sleep(2)
    item = random.randint(0,1000)
    print ("producer notify: producer item number {}".format(item))

    semaphore.release() #+1


if __name__ == '__main__':
    for i in range(0,5):
        t2 = threading.Thread(target=consumer)
        t1 = threading.Thread(target=producer)
        t1.start()
        t2.start()
        t1.join()
        t2.join()
    print ("program terminated")

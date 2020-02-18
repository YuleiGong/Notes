#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 15:02:34
from __future__ import unicode_literals
from __future__ import absolute_import

import threading
import time

class Box(object):
    lock = threading.RLock()

    def __init__(self):
        self.total_items = 0

    def execute(self,n):
        Box.lock.acquire()
        self.total_items += n
        Box.lock.release()

    def add(self):
        Box.lock.acquire()
        self.execute(1)
        Box.lock.release()

    def remove(self):
        Box.lock.acquire()
        self.execute(-1)
        Box.lock.release()


def adder(box,items):
    while items > 0:
        print ("adding 1 item in box")
        box.add()
        time.sleep(1)
        items -= 1

def remover(box, items):
    while items:
        print ("removing 1 item in box")
        box.remove()
        time.sleep(1)
        items -= 1

if __name__ == '__main__':
    items = 5
    print ("putting {} items in the box".format(items))
    box = Box()
    t1 = threading.Thread(target=adder, args=(box, items))
    t2 = threading.Thread(target=remover, args=(box, items))
    t1.start()
    t2.start()

    t1.join()
    t2.join()

    print ("{} items still remain in the box".format(box.total_items))



